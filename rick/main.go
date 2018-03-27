package main

import (
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"sync"
	"log"
	"time"
	"net/http"
	"flag"
	gmp "github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcrowley/go-metrics"
)

var metricsUri *string
var metricsPort *string

func main() {

	brokerUrl := flag.String("broker", "localhost:9092", "broker endpoint in the form of address:port")
	topic := flag.String("topic", "test", "name of the topic")
	message := flag.String("message", "test", "data to send")
	metricsUri = flag.String("metrics-uri", "metrics", "name of the prometheus metrics endpoint")
	metricsPort = flag.String("metrics-port", "8080", "port of the http listener where prometheus metrics will be exposed")

	flag.Parse()

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{*brokerUrl}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)

	var (
		wg                sync.WaitGroup
		successes, errors int
	)

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for range producer.Successes() {
	//		successes++
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for err := range producer.Errors() {
	//		log.Println(err)
	//		errors++
	//	}
	//}()

	setupMetrics(config.MetricRegistry)

	start := currentTimestamp()

//ProducerLoop:
	for  {
		message := &sarama.ProducerMessage{Topic: *topic, Value: sarama.StringEncoder(*message)}
		producer.SendMessage(message)
		//select {
		//case producer.Input() <- message:
		//
		//case <-interrupted:
		//	producer.AsyncClose()
		//	break ProducerLoop
		//}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d in %dms\n", successes, errors, currentTimestamp()-start)
}

func setupMetrics(metricsRegistry metrics.Registry) {
	prometheusRegistry := prometheus.DefaultRegisterer
	pClient := gmp.NewPrometheusProvider(metricsRegistry, "rickandmorty", "rick", prometheusRegistry, 1*time.Second)
	go pClient.UpdatePrometheusMetrics()

	http.Handle("/" + *metricsUri, promhttp.Handler())
	go http.ListenAndServe(":" + *metricsPort, nil)
}

func currentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
