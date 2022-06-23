package version

import (
	"context"
	"fmt"
	"ginco/framework/contract"
	"ginco/framework/elasticsearch"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
	"reflect"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func Command(a contract.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get Application version",
		Long:  "Get Application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("The Application version is v%s\n", a.Version())

			ctx := context.Background()
			es := a.GetI("es").(*elasticsearch.Elasticsearch)

			info, code, err := es.Ping("http://192.168.17.129:9200").Do(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

			query := elastic.NewMatchAllQuery()
			resp, err := es.Search().Index("user").Query(query).Do(ctx)
			if err != nil {
				panic(err)
			}

			var person Person
			var persons []Person
			for _, item := range resp.Each(reflect.TypeOf(person)) {
				if t, ok := item.(Person); ok {
					persons = append(persons, t)
				}
			}

			fmt.Printf("persons:\n %+#v\n", persons)

		},
	}
}
