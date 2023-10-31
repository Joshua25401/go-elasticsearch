package elastic

import (
	"context"
	"elasticsearch-learn/model"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticDependencies struct{}

type elasticSearch struct {
	client *elasticsearch.Client
}

func NewElasticClient(config elasticsearch.Config, dep ElasticDependencies) SearchEngines {
	elasticClient, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatal("cannot instantiate elastic client because of ", err)
	}

	return &elasticSearch{
		client: elasticClient,
	}
}

func (ec *elasticSearch) Search(ctx context.Context, keySearch string) ([]model.User, error) {
	return []model.User{}, nil
}

func (ec *elasticSearch) SearchByType(ctx context.Context, keySearch string, searchType model.UserSearchType) ([]model.User, error) {
	return []model.User{}, nil
}

func (ec *elasticSearch) InsertData(ctx context.Context, newUser model.User) error {
	return nil
}

func (ec *elasticSearch) UpdateData(ctx context.Context, updatedUser model.User) error {
	return nil
}

func (ec *elasticSearch) DeleteData(ctx context.Context, deleteTarget model.User) error {
	return nil
}

func (ec *elasticSearch) SearchEngineInfo(ctx context.Context) {
	res, err := ec.client.Info(ec.client.Info.WithContext(ctx))
	if err != nil {
		log.Fatal("unable to get elasticsearch information because of ", err)
	}
	fmt.Println(res)
}
