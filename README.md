# modbus-go
import : https://github.com/goburrow/modbus

# Usage
```go
mem := memory.New(memory.FUNC_READ_HOLDING_REGISTERS, 0, parser.BIG_LOWER, parser.INT16ARR, 1)
mb := NewTCP(&client.TCP{
  Address: "127.0.0.1:502",
  SlaveID: 1,
  Timeout: 60 * time.Second,
}, mem)
if err := mb.Connect(); err != nil {
  fmt.Println(err)
  return
} else {
  defer mb.Close()
}

chunk := mb.Read()
for i := range chunk {
  fmt.Println("%+v", chunk[i])
}
```
