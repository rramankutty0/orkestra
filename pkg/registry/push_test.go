package registry

import (
	"github.com/go-logr/logr"
	"helm.sh/helm/v3/pkg/chart"
	"net/http"
	"reflect"
	"testing"
)

func TestPushToStaging(t *testing.T) {
	type args struct {
		ch             *chart.Chart
		logr           logr.Logger
		username       string
		password       string
		artifactoryURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushToStaging(tt.args.ch, tt.args.logr, tt.args.username, tt.args.password, tt.args.artifactoryURL); (err != nil) != tt.wantErr {
				t.Errorf("PushToStaging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pushChart(t *testing.T) {
	type args struct {
		ch       *chart.Chart
		username string
		password string
		repoURL  string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pushChart(tt.args.ch, tt.args.username, tt.args.password, tt.args.repoURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("pushChart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pushChart() got = %v, want %v", got, tt.want)
			}
		})
	}
}