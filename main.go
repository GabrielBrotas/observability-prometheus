package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/*
	Gauge Metric => variable values 1, 100, 4, 23, 17,...
	This metric will be responsible for check the online users
*/
var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "goapp_online_users", // identifier
	Help: "Online users", // description
	ConstLabels: map[string]string{
		"course": "fullcycle", // tags
	},
})

/*
	total of http request that was made, incremental value
	1 - 10
	2 - 15 
	3 - 16
	from value 1 to 2 was executed 5 requests, from 2 - 3 just one request
*/
var httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "goapp_http_requests_total", // identifier 
	Help: "Count of all HTTP requests for goapp", // description
}, []string{})

/*

*/
var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "goapp_http_request_duration",
	Help: "Duration in seconds of all HTTP requests",
	// handler is the tag we will use on the MustCurryWith function
}, []string{"handler"})

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers) // registry our metric
	r.MustRegister(httpRequestsTotal) // registry metric
	r.MustRegister(httpDuration)

	// this function will be executed in another thread
	go func() {
		for {
			// infinity loop to change the online users all the time from 1 - 2000
			onlineUsers.Set(float64(rand.Intn(2000)))
		}
	}()

	// function to handle the root function
	home_response := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(8))*time.Second) // mock to the page take longer to execute
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Full Cycle"))
	})

	contact_response := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(5))*time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Contact"))
	})


	// InstrumentHandlerDuration => how long it's taking to execute this function
	home_with_duration := promhttp.InstrumentHandlerDuration(
		httpDuration.MustCurryWith(prometheus.Labels{"handler": "home"}),
		// function to increase the counter on the / page
		promhttp.InstrumentHandlerCounter(httpRequestsTotal, home_response),
	)

	contact_with_duration := promhttp.InstrumentHandlerDuration(
		httpDuration.MustCurryWith(prometheus.Labels{"handler": "contact"}),
		// function to increase the counter on the /contact page
		promhttp.InstrumentHandlerCounter(httpRequestsTotal, contact_response),
	)


	http.Handle("/", home_with_duration)
	http.Handle("/contact", contact_with_duration)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{})) // route where the metrics will be available for prometheus
	log.Fatal(http.ListenAndServe(":8181", nil)) // initializate app
}
