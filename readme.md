## 依赖结构

```bash
windranger proto --target=internal --go_package=github.com/wzyjerry/mahjong internal/util/*/*.yml
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/ent/schema/*/*.proto
wire ./...
```
