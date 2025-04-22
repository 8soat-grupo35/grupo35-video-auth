aws eks update-kubeconfig --name grupo35-video-processing --region us-east-1

sh infra/kubernetes_down.sh

USER_POOL_ID="us-east-1_jJ85wjTux"
APP_CLIENT_ID="2k372ncttsj05lud5v38do9b6p"
AWS_REGION="us-east-1" # Ex: us-east-1
AWS_ACCOUNT_ID="599807457469"

CLIENT_SECRET=$(aws cognito-idp describe-user-pool-client --user-pool-id "$USER_POOL_ID" --client-id "$APP_CLIENT_ID" --query 'ClientSecret' --output text)
REDIRECT_URL="https://example.com/callback"

ISSUER_URL="https://cognito-idp.${AWS_REGION}.amazonaws.com/${USER_POOL_ID}"

echo "Client ID: $APP_CLIENT_ID"
echo "Client Secret: $CLIENT_SECRET"
echo "Redirect URL: $REDIRECT_URL"
echo "Issuer URL: $ISSUER_URL"

docker build -t video-auth:latest .

aws ecr describe-repositories --repository-names video-auth || aws ecr create-repository --repository-name video-auth

aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
docker tag video-auth:latest $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/video-auth:latest
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/video-auth:latest


kubectl create secret generic cognito-secrets \
    --from-literal=cognito-client-id="$APP_CLIENT_ID" \
    --from-literal=cognito-client-secret="$CLIENT_SECRET" \
    --from-literal=cognito-redirect-url="$REDIRECT_URL" \
    --from-literal=cognito-issuer-url="$ISSUER_URL"  --dry-run=client -o yaml | kubectl apply -f -

envsubst < infra/deployment.yaml | kubectl apply -f -
          kubectl apply -f infra/service.yaml
          kubectl set image deployment/video-auth video-auth=$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/video-auth:latest