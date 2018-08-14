package main

import (
	"github.com/argoproj/argo-events/common"
	sv1 "github.com/argoproj/argo-events/pkg/sensor-client/clientset/versioned/typed/sensor/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"github.com/rs/zerolog"
	sc "github.com/argoproj/argo-events/sensor-controller"
)

func main() {
	kubeConfig, _ := os.LookupEnv(common.EnvVarKubeConfig)
	restConfig, err := common.GetClientConfig(kubeConfig)
	if err != nil {
		panic(err)
	}

	sensorName, _ := os.LookupEnv(common.SensorName)
	if sensorName == "" {
		panic("sensor name is not provided")
	}

	sensorNamespace, _ := os.LookupEnv(common.SensorNamespace)
	if sensorNamespace == "" {
		panic("sensor namespace is not provided")
	}

	serverAddr, _ := os.LookupEnv(common.DefaultSensorControllerDeploymentName)
	if serverAddr == "" {
		panic("sensor controller name not provided")
	}

	log := zerolog.New(os.Stdout).With().Str("sensor-name", sensorName).Logger()

	sensorClient, err := sv1.NewForConfig(restConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get sensor client")
	}
	kubeClient := kubernetes.NewForConfigOrDie(restConfig)

	sensor, err := sensorClient.Sensors(sensorNamespace).Get(sensorName, metav1.GetOptions{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to retrieve sensor")
	}

	sCtx := sc.NewSensorContext(sensorClient, kubeClient, restConfig, sensor, log)
	sCtx.StartNotificationHandler()
}
