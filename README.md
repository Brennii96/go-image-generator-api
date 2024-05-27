## Go Image Generator API

I've built a very basic Go API which generates an image from a given prompt. It uses OpenAI's Dalle-3 model so speed of response is dependent on that. By default it runs on port 8080.

To get running clone the repo: 
```
git clone https://github.com/Brennii96/go-image-generator-api.git
```
cd into the directory:
```
cd go-image-generator-api
```
Install the dependencies
```
go get .
```
Generate an auth token:
```
openssl rand -base64 32
```

Get an API key from Open AI https://platform.openai.com/api-keys.
Now put both the auth token and API Key into a .env file:
```
cat <<EOF > .env
OPENAI_API_KEY=your_api_key_here
AUTH_TOKEN=your_auth_token_here
EOF
```
Now everything is setup you should be good to run the app and start calling the API:
```
go run main.go
```

There is only 1 endpoint for generating the image "/api/generate/image" which accepts a prompt from JSON. 
It can be called however you like. For testing I used postman, for speed and ease I'll just provide a curl example below. The bearer token will be the Auth token generated earlier and stored in the .env file. In my example the command generated `/JLl8c0az6BOfMjNbJsYdD1mgDbGHokSMczPRpoHFmI=` so I will use that:
```
curl --location 'http://localhost:8080/api/generate/image' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer /JLl8c0az6BOfMjNbJsYdD1mgDbGHokSMczPRpoHFmI=' \
--data '{
    "prompt": "A cat drinking milk from a saucer"
}'
```
