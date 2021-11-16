# test-skriptforproto

This test project showcases how to call a gRPC service from [Skript](https://github.com/SkriptLang/Skript).

## Uses:
- PaperMC as Minecraft server
- Skript plugin for implementing quick solutions on the Minecraft server
- Skript-Reflect addon for Skript to invoke Java code within skripts
- bufbuild tool to generate Protobuf & gRPC client-server code
- gRPC server implemented in the Go programming language
- Gradle to build fat jar from proto generated Java code (copied into server/plugins/skript-reflect)

## How to use:

1. Start the Minecraft server in terminal A:
```shell
cd server && ./start.sh
```
2. Start the gRPC server in terminal B:
```shell
cd skriptforproto && go run .
```
3. Open Minecraft and join your server listening on `localhost`.
4. Execute the Minecraft commands and watch the logs. DONE! :D:
   - `/hello YourName`
   - `/helloAsyncSection YourName` - same as above but non-blocking
   - `/helloAsyncFunc Bob 1` - same as above
   - `/helloAsyncFunc Robin 0` - zero deadline makes rpc time out immediately
