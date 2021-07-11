package config

import "github.com/spf13/viper"

/*
type Config struct {
	HttpServer        string `mapstructure:"HTTP_SERVER"`
	MqttHost          string `mapstructure:"MQTT_HOST"`
	MqttPort          string `mapstructure:"MQTT_PORT"`
	MqttClientName    string `mapstructure:"MQTT_CLIENT_NAME"`
	InfluxHost        string `mapstructure:"INFLUX_HOST"`
	InfluxPort        string `mapstructure:"INFLUX_PORT"`
	InfluxName        string `mapstructure:"INFLUX_NAME"`
	InfluxMeasurement string `mapstructure:"INFLUX_MEASUREMENT"`
	InfluxToken       string `mapstructure:"INFLUX_TOKEN"`
	InfluxOrg         string `mapstructure:"INFLUX_ORG"`
}
*/

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	return
}
