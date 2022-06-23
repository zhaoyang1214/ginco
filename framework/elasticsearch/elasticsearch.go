package elasticsearch

import (
	"ginco/framework/contract"
	"github.com/olivere/elastic/v7"
)

type Client = elastic.Client

type Elasticsearch struct {
	*Client
	app         contract.Application
	connections map[string]*Client
}

var _ contract.Elasticsearch = (*Elasticsearch)(nil)

func NewElasticsearch(app contract.Application) *Elasticsearch {
	es := &Elasticsearch{
		app:         app,
		connections: make(map[string]*Client),
	}
	es.Client = es.Connection()
	return es
}

func (e *Elasticsearch) Connection(name ...string) *Client {
	var key string
	if len(name) > 0 {
		key = name[0]
	} else {
		key = e.app.GetI("config").(contract.Config).GetString("elasticsearch.default")
	}

	if c, ok := e.connections[key]; ok {
		return c
	}

	e.connections[key] = e.Resolve(key)
	return e.connections[key]
}

func (e *Elasticsearch) Resolve(name string) *Client {
	conf := e.app.GetI("config").(contract.Config).Sub("elasticsearch.connections." + name)
	if conf == nil {
		panic("Elasticsearch config [" + name + "] is not defined")
	}

	var options []elastic.ClientOptionFunc
	urls := conf.GetStringSlice("urls")
	sniff := conf.GetBool("sniff")
	options = append(options, elastic.SetURL(urls...), elastic.SetSniff(sniff))

	username := conf.GetString("username")
	password := conf.GetString("password")
	if username != "" && password != "" {
		options = append(options, elastic.SetBasicAuth(username, password))
	}

	client, err := elastic.NewClient(options...)
	if err != nil {
		panic(err)
	}

	return client
}
