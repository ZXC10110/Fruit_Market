def gv

pipeline {
    agent any

    stages {
        stages("init") {
            // step can allowed us to write script in Jenkinsfile
            steps {
                script {
                    gv = load "script.groovy"
                }
            }
        }

        stages("build") {
            //using groovy script 
            script {
                // call function name "buildApp from script.groovy"
                gv.buildApp()
            }
        }
    }
}