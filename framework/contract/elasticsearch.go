package contract

import (
	"github.com/olivere/elastic/v7"
)

type Elasticsearch interface {
	Connection(names ...string) *elastic.Client
	Resolve(name string) *elastic.Client
}
