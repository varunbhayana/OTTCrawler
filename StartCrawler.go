package main

import (
	"OTTCrawler/enum"
	"fmt"
)

func main() {

	fmt.Println(enum.ZEE_5)
}

// func main() {

// // 	import { getQueueManager } from "../queue/manager";
// // import {connect, disconnect} from "../clients/database";
// // import getCacheManager from '../clients/cache';
// // import { loadMaps } from "../helpers/maps";
// // import { connectCelery } from "../clients/celery";

// // const startCrawler = async () => {
// //     process.on('SIGTERM', async function() {
// //         console.log("\nShutting down gracefully\n")
// //         await queueManager.crawlerQueue.close();
// //         await disconnect();
// //         await cacheClient.disconnect();
// //         process.exit(0)
// //     });

// //     process.on('SIGINT', async function () {
// //         console.log("\nShutting down gracefully\n")
// //         await queueManager.crawlerQueue.close();
// //         await disconnect();
// //         await cacheClient.disconnect();
// //         process.exit(0)
// //     });

// //     await connect();
// //     const cacheClient = await getCacheManager();
// //     await cacheClient.connect();
// //     await loadMaps();
// //     await connectCelery();
// //     const queueManager = getQueueManager();
// //     queueManager.initializePrimaryQueue();
// //     queueManager.initializeCrawlerQueue();
// // };

// // export default startCrawler;

// // const { EventEmitter } = require("events");

// // EventEmitter.defaultMaxListeners = 50;

// // if(process.env.NODE_ENV !== "production"){
// //     require('../../babel-register');
// // }

// const startCrawler = require('../workers/crawler').default;
// startCrawler();
// }
