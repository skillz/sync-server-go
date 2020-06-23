# Sync Server
## Build (WIP)
1. Download source and build
	Note: Please follow instructions found in the build folder README.md
2. The build creates the nakama executable. To run the server execute: ./nakama from the command line
3. Please check out the simple client example to see how to interface with the server:
https://github.com/aaron-skillz/sync-server-unity-example
4. Refer to the Nakama documentation on how to create custom modules which can handle game logic:
https://heroiclabs.com/docs/runtime-code-function-reference/#nakama-module
5. Check out the Nakama docs on how to work with the Unity Client:
https://heroiclabs.com/docs/unity-client-guide/
TODO:
Configure SSL
Use external match maker ID and token (which will be provided by Skillz SDK)
heroiclabs.comheroiclabs.com
.NET/Unity client guide - Nakama server
Documentation for the Nakama realtime and social server for games and apps.

## Docker
Docker Hub: https://hub.docker.com/r/skillzint/nakama

### Start Docker Contianer
```
docker-compose -f ./docker-compose.yml up
```

### Stop Docker Continaer
```
docker-compose -f ./docker-compose.yml down
```
