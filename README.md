# grupo35-video-auth
Projeto que utiliza autenticação do Cognito da AWS.


docker build -t irlan/grupo35-video-auth:latest .

docker push irlan/grupo35-video-auth:latest

kubectl apply -f infra/deployment.yaml
kubectl apply -f infra/service.yaml
