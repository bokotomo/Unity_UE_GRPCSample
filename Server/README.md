# grpc-go

GRPC の実行サンプル

## 動かし方

```
cd ./infrastructure/docker
docker-compose up -d
docker-compose exec go-grpc bash
```

## unaryr RPC

２つのターミナルで、server と client を実行  
２つの数字を送ったら加算するやつを実装

```
go run ./app/unaryrpc/server/
go run ./app/unaryrpc/client/
```

## Bidirectional Streaming RPC

２つのターミナルで、server と client を実行  
チャットっぽいの実装

```
go run ./app/bidirectional/server/
go run ./app/bidirectional/client/
```

## create pb

```
// unary rpc 足算
make protoc OUT=./pb/calc NAME=calc.proto

// server streaming rpc 通知
make protoc OUT=./pb/notification NAME=notification.proto

// client streaming rpc アップロード
make protoc OUT=./pb/upload NAME=upload.proto

// client streaming rpc チャット
make protoc OUT=./pb/chat NAME=chat.proto
```
