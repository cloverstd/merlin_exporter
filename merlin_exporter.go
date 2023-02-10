package main

import (
	"flag"
	"github.com/cloverstd/merlin_exporter/merlin_client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	host := flag.String("merlin-host", "", "merlin router host")
	username := flag.String("user", "", "merlin router manage username")
	password := flag.String("pass", "", "merlin router manage password")

	flag.Parse()

	client, err := merlin_client.New(*host, *username, *password)
	if err != nil {
		log.Fatalf("create merlin client failed, %v\n", err)
	}

	bootTimeGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "merlin_boot_time_seconds",
		Help: "Node boot time, in unix timestamp",
	})

	memoryGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "merlin_memory_info",
		Help: "Merlin router memory info",
	}, []string{"type"})

	cpuGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "merlin_cpu_info",
		Help: "Merlin router cpu info",
	}, []string{"number"})

	temperatureGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "merlin_temperature_gauge",
		Help: "Merlin router chip temperature",
	}, []string{"name"})

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(bootTimeGauge, memoryGauge, cpuGauge, temperatureGauge)

	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}))

	go client.Loop(time.Second*5, func(info *merlin_client.Info) {
		bootTimeGauge.Set(float64(info.Uptime.Unix()))
		for key, value := range info.Temperature {
			temperatureGauge.WithLabelValues(key).Set(value)
		}
		memoryGauge.WithLabelValues("total").Set(float64(info.OSInfo.Memory.Total))
		memoryGauge.WithLabelValues("used").Set(float64(info.OSInfo.Memory.Used))
		memoryGauge.WithLabelValues("free").Set(float64(info.OSInfo.Memory.Free))

		for i, cpu := range info.OSInfo.CPU {
			cpuGauge.WithLabelValues(strconv.Itoa(i)).Set(float64(cpu.Usage) * 100 / float64(cpu.Total))
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Merlin Exporter</title></head>
			<body>
			<h1>Node Exporter</h1>
			<p><a href="/metrics">Metrics</a></p>
			</body>
			</html>`))
	})

	svc := &http.Server{
		Handler: mux,
		Addr:    ":9100",
	}
	log.Printf("Merlin router exporter listened on :9100")

	log.Fatalln(svc.ListenAndServe())
}
