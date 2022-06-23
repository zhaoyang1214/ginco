package elasticsearch

import (
	"github.com/zhaoyang1214/ginco/framework/contract"
	"github.com/zhaoyang1214/ginco/framework/elasticsearch"
)

type Elasticsearch struct {
}

func (d *Elasticsearch) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	appServer, err := container.Get("app")
	if err != nil {
		return nil, err
	}

	return elasticsearch.NewElasticsearch(appServer.(contract.Application)), nil
}
