package main

import (
	"flag"
	"fmt"
)

type options struct {
	appName                      string
	domainName                   string
	namespace                    string
	tillerUrl                    string
	helmTLSKeyFile               string
	helmTLSCertificateFile       string
	tillerTLSSkipVerify          bool
	syncPeriod                   int
	installationTimeout          int64
	applicationGatewayImage      string
	applicationGatewayTestsImage string
	eventServiceImage            string
	eventServiceTestsImage       string
}

func parseArgs() *options {
	appName := flag.String("appName", "application-operator", "Name used in controller registration")
	domainName := flag.String("domainName", "kyma.local", "Domain name of the cluster")
	namespace := flag.String("namespace", "kyma-integration", "Namespace in which the Application chart will be installed")
	tillerUrl := flag.String("tillerUrl", "tiller-deploy.kube-system.svc.cluster.local:44134", "Tiller release server url")
	helmTLSKeyFile := flag.String("helmTLSKeyFile", "/etc/certs/tls.key", "Path to TLS key used for Tiller communication")
	helmTLSCertificateFile := flag.String("helmTLSCertificateFile", "/etc/certs/tls.crt", "Path to TLS certificate used for Tiller communication")
	tillerTLSSkipVerify := flag.Bool("tillerTLSSkipVerify", true, "")
	syncPeriod := flag.Int("syncPeriod", 30, "Time period between resyncing existing resources")
	installationTimeout := flag.Int64("installationTimeout", 240, "Time after the release installation will time out")

	applicationGatewayImage := flag.String("applicationGatewayImage", "", "The image of the Application Gateway to use")
	applicationGatewayTestsImage := flag.String("applicationGatewayTestsImage", "", "The image of the Application Gateway Tests to use")
	eventServiceImage := flag.String("eventServiceImage", "", "The image of the Event Service to use")
	eventServiceTestsImage := flag.String("eventServiceTestsImage", "", "The image of the Event Service Tests to use")

	flag.Parse()

	return &options{
		appName:                      *appName,
		domainName:                   *domainName,
		namespace:                    *namespace,
		tillerUrl:                    *tillerUrl,
		helmTLSKeyFile:               *helmTLSKeyFile,
		helmTLSCertificateFile:       *helmTLSCertificateFile,
		tillerTLSSkipVerify:          *tillerTLSSkipVerify,
		syncPeriod:                   *syncPeriod,
		installationTimeout:          *installationTimeout,
		applicationGatewayImage:      *applicationGatewayImage,
		applicationGatewayTestsImage: *applicationGatewayTestsImage,
		eventServiceImage:            *eventServiceImage,
		eventServiceTestsImage:       *eventServiceTestsImage,
	}
}

func (o *options) String() string {
	return fmt.Sprintf("--appName=%s --domainName=%s --namespace=%s --tillerUrl=%s"+
		"--helmTLSKeyFile=%s --helmTLSCertificateFile=%s --tillerTLSSkipVerify=%v --syncPeriod=%d --installationTimeout=%d "+
		"--applicationGatewayImage=%s --applicationGatewayTestsImage=%s --eventServiceImage=%s --eventServiceTestsImage=%s",
		o.appName, o.domainName, o.namespace, o.tillerUrl,
		o.helmTLSKeyFile, o.helmTLSCertificateFile, o.tillerTLSSkipVerify, o.syncPeriod, o.installationTimeout,
		o.applicationGatewayImage, o.applicationGatewayTestsImage, o.eventServiceImage, o.eventServiceTestsImage)
}
