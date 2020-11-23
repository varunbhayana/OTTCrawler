package queue

import (
	"OTTCrawler/enum"
	"github.com/adjust/rmq/v3"
)

var connection rmq.Connection
func init(){
	con, err := rmq.OpenConnection("my service", "tcp", "localhost:6379", 1,nil)
	if (err!=nil){
		panic(err)
	}
	connection = con
}
type QueueManager struct {
	ID enum.CrawlingIds
	CrawerQueue rmq.Queue
	apiQueues *map[enum.CrawlingIds][rmq.Queue]
}

func (manager *QueueManager) InitializeApiQueue(id enum.CrawlingIds){

}

func (manager *QueueManager) GetApiQueue(id enum.CrawlingIds){
	// taskQueue, err := connection.OpenQueue("tasks")
	//if (!this.apiQueues[id]) {
		if val,ok := manager.apiQueues[id];ok{

		}else{
			apiQueue, err := connection.OpenQueue(id.ToString())
			if (err!=nil){
				panic(err)
			}
			limit := 0
			// if val,ok := conf.D.Queues["ApiRateLimit"];ok{

			// }
			err = apiQueue.StartConsuming(limit, time.Second)
			if (err!=nil){
				panic(err)
			}
			manager.apiQueues[id] = queue
		}
}

