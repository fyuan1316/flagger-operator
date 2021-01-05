package v1alpha1

import "gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/api"

func (in *FlaggerStatus) setState(state api.OperatorState) {
	in.State = state
}

func (in *FlaggerStatus) SetState(isReady, isHealthy bool) {
	if isHealthy {
		in.setState(api.OperatorStates.Health)
		return
	}
	if isReady {
		in.setState(api.OperatorStates.Ready)
		return
	}
	in.setState(api.OperatorStates.NotReady)
	return
}
