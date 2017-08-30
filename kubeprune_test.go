package main

import (
	"reflect"
	"testing"
)

func TestKubePrune(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Kube Namespace",
			input: "context",
			want:  "np",
		},
		{
			name:  "Kube Context",
			input: "namespace",
			want:  "manualsit1",
		},
		{
			name:  "Kube rt",
			input: "rt",
			want:  "ris",
		},
		{
			name:  "Kube domain",
			input: "domain",
			want:  "np",
		},
		{
			name:  "Kube domain",
			input: "environment",
			want:  "kube1",
		},
		{
			name:  "Kube environmentBranchOverride",
			input: "environmentBranchOverride",
			want:  "testbranchoverride",
		},
		{
			name:  "Kube globalBranchOverride",
			input: "globalBranchOverride",
			want:  "testglobalbranchoverride",
		},
		{
			name:  "Kube deploy targets",
			input: "deploy",
			want:  "[\n      \"ipt-idresint/ipt-idresint-svc\",\n      \"ipt-idresorch/ipt-idresorch-svc\",\n      \"ipt-router/ipt-router-svc\"\n    ]",
		},
		{
			name:  "Kube deploy targets",
			input: "exclude",
			want:  "[\n    \"ipt-risid/ipt-risid-bpm\"\n  ]",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := kubeprune(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}
