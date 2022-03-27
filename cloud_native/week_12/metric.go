package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
)

var histogram *prometheus.HistogramVec

func init(){
	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "httpserver",
		Name:      "execution_latency_seconds",
		Help:      "Time spent",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
	}, []string{"step"})

	err := prometheus.Register(histogram)
	if err != nil {
		log.Println(err)
		return
	}
}

func ObserveTotal(start time.Time){
	histogram.WithLabelValues("total").Observe(time.Now().Sub(start).Seconds())
}


