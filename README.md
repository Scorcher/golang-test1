# golang-test1

Console 1
```bash
go run main.go
```

Console 2
```bash
curl http://127.0.0.1:8080/storage/put/key2/value2 &
curl http://127.0.0.1:8080/storage/put/key1/value1 &
curl http://127.0.0.1:8080/storage/get/key1
```

output
```
Put key "key2", value "value2" | Took: 5000 ms | OK
Put key "key1", value "value1" | Took: 8389 ms | OK
Get key "key1" | Took: 0 ms | Value: "value1"
```

В коде в методе Put намеренно выставлен sleep(5s)

На Get сделал `RLock`
