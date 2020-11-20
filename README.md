# Simple DYNDNS Server for Cloudflare

## Environment variables

- `APP_MODE=release`
- `APP_PORT=8080`
- `CLOUDFLARE_API_KEY` => Global API Key (Cloudflare User Settings)
- `CLOUDFLARE_ZONE` => Name of the zone / domain ("test.de")
- `CLOUDFLARE_MAIL` => Your email address used for your Cloudflare user account
- `UPDATE_PASSWORD` => Password used to authorize the update request with the username "update"


## Setup

`go build`

Use a reverse proxy for ssl encryption!

## Docker Setup

Build Docker container:
`docker build -t cloudflare-dyndns:0.1 .`

Run Docker container:
```
docker run --name <container-name> -d \
    --restart=always \
    -p 127.0.0.1:8080:8080 \
    -e APP_MODE="release" \
    -e APP_PORT="8080" \
    -e CLOUDFLARE_API_KEY="<your-cloudflare-key>" \
    -e CLOUDFLARE_ZONE="<your-cloudflare-zone>" \
    -e CLOUDFLARE_MAIL="<your-cloudflare-email>" \
    -e UPDATE_PASSWORD="<password for update user>" \
    cloudflare-dyndns:0.1
```