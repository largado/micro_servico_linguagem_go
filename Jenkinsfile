pipeline {
    agent any

    stages {
        stage ("Começo do processo...") {
            steps {
                echo 'Começo da execução da esteira'
                script {
                    dockerapp = docker.build("avena/crud:${env.BUILD_ID}", '-f ./Dockerfile ./')
                }

            }
        }

        stage ('envio da imagem gerada para o docker hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                        dockerapp.push('latest')
                        dockerapp.push("${env.BUILD_ID}")
                    }
                }
            }

        }
    }
}
