# media
Search and play movies, series through different streaming providers.

## Features
- [x] Play movie/series (sending back url + action)
- [x] Support Plex as a provider
- [ ] Support Netflix as a provider?
- [ ] Support Prime Video as a provider?
- [ ] Support Disney Plus as a provider?

## Prerequisites

- Having ``golang`` installed [instructions](https://golang.org/doc/install)

## Build

```bash
$ go build -o bin/ability cmd/ability/main.go
```

## Run

```bash
$ bin/ability
```

## Requests example

#### Play media on a chromecast device

```bash
$ curl -i -H "Content-Type":"application/json" -X POST http://localhost:4444/resolve -d '{"nlu":{"BestIntent": "PLAY_MOVIE","entities":[{"label":"title","text":"matrix"}]},"device": {"instruments": [{"kind": "chromecast", "actions": ["play", "pause", "play_media"], "name": "salon"}]}}'
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 07 Oct 2022 04:57:37 GMT
Content-Length: 327

{"nlg":{"sentence":"Playing {{ title }} on the chrome cast {{ instrument }}.","params":[{"name":"title","value":"Matrix Resurrections","type":"string"},{"name":"instrument","value":"salon","type":"string"}]},"actions":[{"identifier":"play_media","params":[{"key":"instrument","value":"salon"}]}],"context":{"slot_filling":{}}}
```
