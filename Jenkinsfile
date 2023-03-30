pipeline {
    agent any
    environment {
        imageId = "shorten:${currentBuild.number}"
    }
    stages {
        stage("Test") {
            steps{
                sh "docker build --no-cache --force-rm -t ${imageId} --target test ."
                sh "docker rmi ${imageId}"
            }
        }
        stage("Build") {
            steps{
                sh "docker build --no-cache --force-rm -t ${imageId} ."
                sh "echo build successfully"
            }
        }
        stage("Push") {
            steps{
                sh "echo pushing to repo"
            }
        stage("Clean") {
            steps{
                sh "docker rmi ${imageId}"
            }
        }
    }
}
