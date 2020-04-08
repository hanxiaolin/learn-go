package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v6"
	"hanxiaolin/gin-demo/pkg/setting"
	"net/http"
)

var ESConn *elasticsearch.Client

func setUp() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			setting.ElasticsearchSetting.Host + ":" + setting.ElasticsearchSetting.Port,
		},
		Password: setting.ElasticsearchSetting.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: setting.ElasticsearchSetting.MaxIdle,
		},
	}

	ESConn, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return
	}
}

func Search() {
	// ESConn
}
