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