package main

type RocksDBHandler struct {
}

// func (rock *RocksDBHandler) Set(key, value []byte) error {
//     return nil
// }

func (rock *RocksDBHandler) Ping() (*StatusReply, error) {
    return &StatusReply{"PONG"}, nil
}