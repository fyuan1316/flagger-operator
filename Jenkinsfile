library "alauda-cicd"
def language = "golang"
AlaudaPipeline {
    config = [
        agent: 'golang-1.13',
        folder: '.',
        chart: [
            [
                pipeline: "chart-global-asm",
                project: "asm",
                chart: "global-asm",
                component: "flagger_operator",
            ],
        ],
        scm: [
            credentials: 'alaudabot-bitbucket'
        ],
        docker: [
            repository: "asm/flagger-mesh-operator",
            context: ".",
            dockerfile: "Dockerfile",
            armBuild: true,
        ],
        sonar: [
            binding: "sonarqube",
        ],
    ]
    env = [
        GOPROXY: "https://athens.alauda.cn",
    ]
    steps = [
        [
            name: "Unit test",
            container: language,
            groovy: [
                """
                try {
                sh script: "make test", label: "unit tests..."
                } finally {
                archiveArtifacts 'test.json'
                junit allowEmptyResults: true, testResults: 'pkg/**/*.xml'
                }
                """
            ]
        ],
        [
            name: "Build",
            container: language,
            commands: [
                "make build",
                "make armbuild",
            ]
        ]
    ]
}
