echo 'This will create a challenge file, put it in static and change app.yaml then ./deploy.sh to serve it'
sudo certbot certonly --manual --preferred-challenges http -d repos.errbot.io
# now it should have generated the .pem files, cat them in appengine format

echo 'That goes in "PEM encoded X.509 public key certificate" in https://console.cloud.google.com/appengine/settings/certificates?project=errbot-1127&serviceId=default&versionId=1'
sudo cat /etc/letsencrypt/live/repos.errbot.io/fullchain.pem

echo 'That goes in "Unencrypted PEM encoded RSA private key"'
sudo openssl rsa -inform pem -in /etc/letsencrypt/live/repos.errbot.io/privkey.pem -outform pem

