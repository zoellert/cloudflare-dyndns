# Simple DYNDNS Server for Cloudflare

A simple DynDNS server with an HTTP endpoint for FritzBox router, Securepoint UTM and more to update your private ip-address in your Cloudflare zone.
Currently, only IPv4 addresses are supported.

## Environment variables

|Environment Variable|Description|Example-Value|
|---|---|---|
|`APP_MODE`|Application / log level|`release` \ `debug`|
|`APP_PORT`|Webserver port|`8080`|
|`CLOUDFLARE_API_KEY`|Global API Key (Cloudflare User Settings)|`9e1a9405beaa42158ba179c0eb637651`|
|`CLOUDFLARE_ZONE`|Name of the zone / domain|`your-domain.de`|
|`CLOUDFLARE_MAIL`|Email address used for your Cloudflare user account|`mail@your-domain.de`|
|`UPDATE_PASSWORD`|Password used to authorize the update request with the username "update"|`test123456#`|

## Setup

`go build`

Use a reverse proxy for ssl encryption!

## Docker Setup using Docker-Hub

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
    zoellert/cloudflare-dyndns:latest
```

## Manual Docker Setup

Clone Project.

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

Use a reverse proxy for ssl encryption!

## Router Setup

### Securepoint

| Setting | Description | Example |
|---|---|---|
|Hostname|Choose a subdomain from your Cloudflare zone|`home.your-domain.de`|
|Username|Username for authentication|`update`|
|Password|password defined in UPDATE_PASSWORD environment variable|`test123456#`|
|Server|Address to your instance of this dyndns server|`dyndns-server.your-domain.de`|

### FritzBox
| Setting | Description | Example |
|---|---|---|
|URL|Address to your instance of this dyndns server|`https://dnydns-server.your-domain.de/nic/update?hostname=<domain>&myip=<ipaddr>`|
|Domainname|Choose a subdomain from your Cloudflare zone|`home.your-domain.de`|
|Username|Username for authentication|`update`|
|Password|password defined in UPDATE_PASSWORD environment variable|`test123456#`|