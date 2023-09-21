package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"sales-product-srv/global"
	"sales-product-srv/model"
)

func InitElastic() {
	var err error
	host := global.ServerConfig.ElasticConfig.Host
	port := global.ServerConfig.ElasticConfig.Port

	url := fmt.Sprintf("http://%s:%d", host, port)

	logger := log.New(os.Stdout, "sales-product-srv", log.LstdFlags)

	global.ElasticClient, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	// 新建mapping和index
	exists, err := global.ElasticClient.IndexExists(model.ElasticBanner{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		global.ElasticClient.CreateIndex(model.ElasticBanner{}.GetIndexName()).BodyString(model.ElasticBanner{}.GetMapping()).Do(context.Background())
	}

}
