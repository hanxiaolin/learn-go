//package elasticsearch
//
//import "github.com/elastic/go-elasticsearch/v6"
//
//var ESConn *elasticsearch.Client
//
//func setUp() {
//	cfg := elasticsearch.Config{
//		Addresses: []string{
//			"http://localhost:9201",
//		},
//		// ...
//	}
//	ESConn, err := elasticsearch.NewClient(cfg)
//	if err != nil {
//		return
//	}
//}
