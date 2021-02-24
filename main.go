/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"github.com/fyuan1316/flagger-operator/pkg/task/entry"
	"github.com/fyuan1316/flagger-operator/pkg/util/env"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	operatorv1alpha1 "github.com/fyuan1316/flagger-operator/api/v1alpha1"
	"github.com/fyuan1316/flagger-operator/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(operatorv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var interval int
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.IntVar(&interval, "interval", 30, "Timer interval is used to synchronize business cluster resources")
	flag.Parse()
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
	watchNS, err := env.GetWatchNamespace()
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}
	leaderElectionNS, leaderElectionEnabled := env.GetLeaderElectionNamespace()
	d := time.Second * time.Duration(interval)
	if !leaderElectionEnabled {
		setupLog.Info("Leader election namespace not set. Leader election is disabled. NOT APPROPRIATE FOR PRODUCTION USE!")
	}
	var mgrOpt manager.Options
	if watchNS != "" {
		namespaces := strings.Split(watchNS, ",")
		// Create MultiNamespacedCache with watched namespaces if it's not empty.
		mgrOpt = manager.Options{
			NewCache:                cache.MultiNamespacedCacheBuilder(namespaces),
			Scheme:                  scheme,
			MetricsBindAddress:      metricsAddr,
			Port:                    9443,
			LeaderElection:          leaderElectionEnabled,
			LeaderElectionNamespace: leaderElectionNS,
			LeaderElectionID:        "flagger-operator-lock",
			SyncPeriod:              &d,
		}
	} else {
		// Create manager option for watching all namespaces.
		mgrOpt = manager.Options{
			Scheme:                  scheme,
			MetricsBindAddress:      metricsAddr,
			Port:                    9443,
			Namespace:               watchNS,
			LeaderElection:          leaderElectionEnabled,
			LeaderElectionNamespace: leaderElectionNS,
			LeaderElectionID:        "flagger-operator-lock",
			SyncPeriod:              &d,
		}
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), mgrOpt)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.FlaggerReconciler{
		Client:   mgr.GetClient(),
		Log:      ctrl.Log.WithName("flagger-operator").WithName("Flagger"),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor("flagger-operator"),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Flagger")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	//setup tasks
	if err := entry.SetUp(); err != nil {
		panic(err)
	}
	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
