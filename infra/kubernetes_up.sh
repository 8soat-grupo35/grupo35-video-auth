sh infra/kubernetes_down.sh

docker build -t video-auth:latest .

aws ecr describe-repositories --repository-names video-auth || aws ecr create-repository --repository-name video-auth

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 599807457469.dkr.ecr.us-east-1.amazonaws.com
          docker tag video-auth:latest 599807457469.dkr.ecr.us-east-1.amazonaws.com/video-auth:latest
          docker push 599807457469.dkr.ecr.us-east-1.amazonaws.com/video-auth:latest

envsubst < infra/deployment.yaml | kubectl apply -f -
          kubectl apply -f infra/service.yaml
          kubectl set image deployment/video-auth video-auth=599807457469.dkr.ecr.us-east-1.amazonaws.com/video-auth:latest