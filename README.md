# API gateway for [Silverbullet](https://silverbullet.md/) 

With the help of [silverbullet-go-api](https://github.com/Mrton0121/silverbullet-go-api) I created this little service.

It runs in a docker container and interacts with Silverbullet API to add data to a page with some ✨magic variables✨.

You can define a pattern in which the data will be sent to the SB API and it will create/append the page you choose.

## Installation:

You just need to copy the following `docker-compose.yml` to your host and overwrite the environment variables and hit `docker compose up -d`.

```
services:
  sb-api-gateway:
    image: mrton0121/silverbullet-api-gateway
    restart: unless-stopped
    environment:
    - SB_URL=http://your-host.xyz:3000 # REQUIRED, The url of your Silverbullet instance
    - SB_TOKEN=your-api-token # REQUIRED, your SB_AUTH_TOKEN env variable
    - SB_PAGE=the-page-you-want-to-append.md # REQUIRED, the page you want to paste the data
    - DATA_PATTERN=**[TEXT]** #defaults to only the text you send in the request
    - SEPARATOR=moses # OPTIONAL, defaults to \n
    ports:
      - 8080:8080
```

## Magic variables:

- **[TEXT]**: the data read from the POST request
- **[DATE]**: the current date in timedate format ( e.g: `2006-01-02 15:04:05` )
- **[SEPARATOR]**: the separator string you defined in the ENV variable
- **[TAB]**: indentation ( `\t` char )

## Example usage: 

```
curl -X POST -d "data=HelloWorld!" http://localhost:8080
```


## Use cases:

I made this service because I wanted a quick way for saving my clipboard on my iPhone. I created a shortcut on my phone which sends a POST request to this service and it gets written into silverbullet. This way I can quickly share urls, texts, almost everything. (Pictures are not supported in the Shortcuts Clipboard variable AFAIK)

I'm sure you can come up with something too if you're interested in it. Tickets are welcome!


Author: Márton Gombócz (Mrton0121)