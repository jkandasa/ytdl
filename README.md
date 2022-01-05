# ytdl
This tool is based on https://github.com/kkdai/youtube

**This tool is meant to be used to download CC0 licenced content, we do not support nor recommend using it for illegal activities.**

### Installation
```bash
docker run --detach --name ytdl \
    --publish 8080:8080 \
    --env TZ="Asia/Kolkata" \
    --restart unless-stopped \
    docker.io/jkandasa/ytdl:master
```