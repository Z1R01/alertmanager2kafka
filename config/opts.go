package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type (
	Opts struct {
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		Kafka struct {
			Host          string `long:"kafka.host"                 env:"KAFKA_HOST"                        description:"Kafka host, eg. kafka-0:9092" required:"true"`
			Topic         string `long:"kafka.topic"                env:"KAFKA_TOPIC"                       description:"Kafka topic, eg. alertmanager" required:"true"`
		SSLCert       string `long:"kafka.ssl.cert"                env:"KAFKA_SSL_CERT" env:"KAFKA_TLS_CERT"                       description:"Kafka client SSL certificate file" required:"false"`
		SSLKey        string `long:"kafka.ssl.key"                env:"KAFKA_SSL_KEY" env:"KAFKA_TLS_KEY"                       description:"Kafka client SSL key file" required:"false"`
		SSLCACert     string `long:"kafka.ssl.cacert"                env:"KAFKA_SSL_CACERT" env:"KAFKA_TLS_CA"                       description:"Kafka server CA certificate file" required:"false"`
			SASLUsername  string `long:"kafka.sasl.username"            env:"KAFKA_SASL_USERNAME"                description:"Kafka SASL username" required:"false"`
			SASLPassword  string `long:"kafka.sasl.password"            env:"KAFKA_SASL_PASSWORD"                description:"Kafka SASL password" required:"false"`
			SASLMechanism string `long:"kafka.sasl.mechanism"           env:"KAFKA_SASL_MECHANISM"               description:"Kafka SASL mechanism (PLAIN, SCRAM-SHA-256, SCRAM-SHA-512)" required:"false" default:"SCRAM-SHA-512"`
		}

		ServerBind string `long:"bind"     env:"SERVER_BIND"   description:"Server address"     default:":9097"`
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}

