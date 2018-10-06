package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/berfinsari/envoyproxybeat/config"
)

// Envoyproxybeat configuration.
type Envoyproxybeat struct {
	done     chan struct{}
	config   config.EnvoyproxybeatConfig
	client   beat.Client
	EbConfig config.ConfigSettings
	period   time.Duration
	port     string
	host     string
}

const selector = "envoyproxybeat"

// New creates an instance of envoyproxybeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	eb := &Envoyproxybeat{
		done: make(chan struct{}),
	}
	err := cfgfile.Read(&eb.EbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return nil, fmt.Errorf("Error reading configuration file: %v", err)
	}
	return eb, nil
}

// Run starts envoyproxybeat.
func (eb *Envoyproxybeat) Run(b *beat.Beat) error {
	logp.Info("envoyproxybeat is running! Hit CTRL-C to stop it.")
	eb.CheckConfig(b)

	var err error
	eb.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(eb.period)
	for {
		select {
		case <-eb.done:
			return nil
		case <-ticker.C:
		}
		serverstats, err := eb.getServerStats(b)
		if err != nil {
			logp.Debug(selector, "Error reading server stats")
			return err
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":   b.Info.Name,
				"server": serverstats,
			},
		}
		eb.client.Publish(event)
		logp.Info("Event sent")
	}
}

func (eb *Envoyproxybeat) CheckConfig(b *beat.Beat) error {
	if eb.EbConfig.Input.Period != nil {
		eb.period = time.Duration(*eb.EbConfig.Input.Period) * time.Second
	} else {
		eb.period = 30 * time.Second
	}

	if eb.EbConfig.Input.Port != nil {
		eb.port = *eb.EbConfig.Input.Port
	} else {
		eb.port = "9901"
	}

	if eb.EbConfig.Input.Host != nil {
		eb.host = *eb.EbConfig.Input.Host
	} else {
		eb.port = "localhost"
	}

	logp.Debug(selector, "Init Envoyproxybeat")
	logp.Debug(selector, "Port %v", eb.port)
	logp.Debug(selector, "Period %v", eb.period)
	logp.Debug(selector, "Host %v", eb.host)

	return nil
}

// Stop stops envoyproxybeat.
func (eb *Envoyproxybeat) Stop() {
	eb.client.Close()
	close(eb.done)
}
