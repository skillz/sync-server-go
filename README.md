# Sync Server
## Build (WIP)
1. Download source and build
	Note: Please follow instructions found in the build folder README.md
2. The build creates the nakama executable. To run the server execute: ./nakama from the command line
3. Check out the Nakama docs on how to work with the Unity Client:
https://heroiclabs.com/docs/unity-client-guide/
4. Refer to the Nakama documentation on how to create custom modules which can handle game logic:
https://heroiclabs.com/docs/runtime-code-function-reference/#nakama-module
5. Please check out the simple client example to see how to interface with the server:
https://github.com/aaron-skillz/sync-server-unity-example


## Pre-Build Server
We have created a docker image that you can use. The server uses the Nakama interface and API so you can use the examples to work with the server. You can download the image and run it locally.

## Docker Image
Docker Hub: https://hub.docker.com/r/skillzint/nakama

### Start Docker Contianer
```
docker-compose -f ./docker-compose.yml up
```

### Stop Docker Continaer
```
docker-compose -f ./docker-compose.yml down
```
