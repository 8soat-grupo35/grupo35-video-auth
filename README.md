# grupo35-video-auth
Projeto que utiliza autenticação do Cognito da AWS.


docker build -t irlan/grupo35-video-auth:latest .

docker push irlan/grupo35-video-auth:latest

kubectl apply -f infra/deployment.yaml
kubectl apply -f infra/service.yaml

aws eks update-kubeconfig --name grupo35-video-processing --region us-east-1


Run tests

Para rodar os mocks
go generate ./...


Para adicionar uma interface no gerador de mocks
//go:generate mockgen -source={nome do arquivo}.go -destination=mock/{nome do arquivo}.go