# get-video-duration
Go binary to get the duration of a video in seconds.

## Deployment using docker
```docker run -p 8080:8080 codehimanshu/get-video-duration```

## Install from source
Prerequisites:
- go 1.13+
- ffmpeg

Start API Server:
- go run main.go

## Usage
Curl Example:
```
curl http://localhost:8080/duration?url=http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4
```
Response
```
{"duration":596.473333}
```

## ToDos
- Use environment variables to get Port and path to ffmpeg/ffprobe

