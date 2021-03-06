// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package spark

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "spark-history-server-url"), defaultConfig.SparkHistoryServerURL, "URL for SparkHistory Server that each job will publish the execution history to.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.mixed.cloudwatch-enabled"), defaultConfig.LogConfig.Mixed.IsCloudwatchEnabled, "Enable Cloudwatch Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.cloudwatch-region"), defaultConfig.LogConfig.Mixed.CloudwatchRegion, "AWS region in which Cloudwatch logs are stored.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.cloudwatch-log-group"), defaultConfig.LogConfig.Mixed.CloudwatchLogGroup, "Log group to which streams are associated.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.cloudwatch-template-uri"), defaultConfig.LogConfig.Mixed.CloudwatchTemplateURI, "Template Uri to use when building cloudwatch log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.mixed.kubernetes-enabled"), defaultConfig.LogConfig.Mixed.IsKubernetesEnabled, "Enable Kubernetes Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.kubernetes-url"), defaultConfig.LogConfig.Mixed.KubernetesURL, "Console URL for Kubernetes logs")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.kubernetes-template-uri"), defaultConfig.LogConfig.Mixed.KubernetesTemplateURI, "Template Uri to use when building kubernetes log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.mixed.stackdriver-enabled"), defaultConfig.LogConfig.Mixed.IsStackDriverEnabled, "Enable Log-links to stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.gcp-project"), defaultConfig.LogConfig.Mixed.GCPProjectName, "Name of the project in GCP")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.stackdriver-logresourcename"), defaultConfig.LogConfig.Mixed.StackdriverLogResourceName, "Name of the logresource in stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.mixed.stackdriver-template-uri"), defaultConfig.LogConfig.Mixed.StackDriverTemplateURI, "Template Uri to use when building stackdriver log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.user.cloudwatch-enabled"), defaultConfig.LogConfig.User.IsCloudwatchEnabled, "Enable Cloudwatch Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.cloudwatch-region"), defaultConfig.LogConfig.User.CloudwatchRegion, "AWS region in which Cloudwatch logs are stored.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.cloudwatch-log-group"), defaultConfig.LogConfig.User.CloudwatchLogGroup, "Log group to which streams are associated.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.cloudwatch-template-uri"), defaultConfig.LogConfig.User.CloudwatchTemplateURI, "Template Uri to use when building cloudwatch log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.user.kubernetes-enabled"), defaultConfig.LogConfig.User.IsKubernetesEnabled, "Enable Kubernetes Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.kubernetes-url"), defaultConfig.LogConfig.User.KubernetesURL, "Console URL for Kubernetes logs")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.kubernetes-template-uri"), defaultConfig.LogConfig.User.KubernetesTemplateURI, "Template Uri to use when building kubernetes log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.user.stackdriver-enabled"), defaultConfig.LogConfig.User.IsStackDriverEnabled, "Enable Log-links to stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.gcp-project"), defaultConfig.LogConfig.User.GCPProjectName, "Name of the project in GCP")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.stackdriver-logresourcename"), defaultConfig.LogConfig.User.StackdriverLogResourceName, "Name of the logresource in stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.user.stackdriver-template-uri"), defaultConfig.LogConfig.User.StackDriverTemplateURI, "Template Uri to use when building stackdriver log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.system.cloudwatch-enabled"), defaultConfig.LogConfig.System.IsCloudwatchEnabled, "Enable Cloudwatch Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.cloudwatch-region"), defaultConfig.LogConfig.System.CloudwatchRegion, "AWS region in which Cloudwatch logs are stored.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.cloudwatch-log-group"), defaultConfig.LogConfig.System.CloudwatchLogGroup, "Log group to which streams are associated.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.cloudwatch-template-uri"), defaultConfig.LogConfig.System.CloudwatchTemplateURI, "Template Uri to use when building cloudwatch log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.system.kubernetes-enabled"), defaultConfig.LogConfig.System.IsKubernetesEnabled, "Enable Kubernetes Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.kubernetes-url"), defaultConfig.LogConfig.System.KubernetesURL, "Console URL for Kubernetes logs")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.kubernetes-template-uri"), defaultConfig.LogConfig.System.KubernetesTemplateURI, "Template Uri to use when building kubernetes log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.system.stackdriver-enabled"), defaultConfig.LogConfig.System.IsStackDriverEnabled, "Enable Log-links to stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.gcp-project"), defaultConfig.LogConfig.System.GCPProjectName, "Name of the project in GCP")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.stackdriver-logresourcename"), defaultConfig.LogConfig.System.StackdriverLogResourceName, "Name of the logresource in stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.system.stackdriver-template-uri"), defaultConfig.LogConfig.System.StackDriverTemplateURI, "Template Uri to use when building stackdriver log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.all-user.cloudwatch-enabled"), defaultConfig.LogConfig.AllUser.IsCloudwatchEnabled, "Enable Cloudwatch Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.cloudwatch-region"), defaultConfig.LogConfig.AllUser.CloudwatchRegion, "AWS region in which Cloudwatch logs are stored.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.cloudwatch-log-group"), defaultConfig.LogConfig.AllUser.CloudwatchLogGroup, "Log group to which streams are associated.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.cloudwatch-template-uri"), defaultConfig.LogConfig.AllUser.CloudwatchTemplateURI, "Template Uri to use when building cloudwatch log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.all-user.kubernetes-enabled"), defaultConfig.LogConfig.AllUser.IsKubernetesEnabled, "Enable Kubernetes Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.kubernetes-url"), defaultConfig.LogConfig.AllUser.KubernetesURL, "Console URL for Kubernetes logs")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.kubernetes-template-uri"), defaultConfig.LogConfig.AllUser.KubernetesTemplateURI, "Template Uri to use when building kubernetes log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.all-user.stackdriver-enabled"), defaultConfig.LogConfig.AllUser.IsStackDriverEnabled, "Enable Log-links to stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.gcp-project"), defaultConfig.LogConfig.AllUser.GCPProjectName, "Name of the project in GCP")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.stackdriver-logresourcename"), defaultConfig.LogConfig.AllUser.StackdriverLogResourceName, "Name of the logresource in stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.all-user.stackdriver-template-uri"), defaultConfig.LogConfig.AllUser.StackDriverTemplateURI, "Template Uri to use when building stackdriver log links")
	return cmdFlags
}
