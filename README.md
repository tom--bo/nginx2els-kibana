## Nginxのaccesslogをkibanaに投げつけるやつ



## 雑な操作スクリプト
```
# template指定でindex, typeの作成
curl -XPOST 'localhost:9200/isucon-3' -d @elastic_template.json

# テスト的にサンプルレコードを投げる
curl -XPOST 127.0.0.1:9200/isucon-3/access -d @test.json

# accesslogを作成したindex,typeに投げる
go run ltsv2els.go access.log http://127.0.0.1:9200/isucon-3/access
```


