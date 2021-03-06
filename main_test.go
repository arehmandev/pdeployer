package main

import (
	"reflect"
	"testing"
)

func TestKubefunctionsdebug(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{
			name: "Kube Debug function",
			want: "pdeploy is currently using these settings",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := debug()
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}

func TestKubePrune(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Kube Namespace",
			input: "context",
			want:  "minikube",
		},
		{
			name:  "Kube Context",
			input: "namespace",
			want:  "polyshop",
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

func TestKubecontext(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Kube context flag test",
			input: "",
			want:  "The current context is kube",
		},
		{
			name:  "Kube context flag test",
			input: "testcontext",
			want:  "The current context is testcontext",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := context(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}

func TestKubevertical(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Kube vertical true flag test",
			input: "true",
			want:  "vertical is set to true",
		},
		{
			name:  "Kube vertical false flag test",
			input: "false",
			want:  "vertical is set to false",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := vertical(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}

func TestKubepublish(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{
			name: "Kube publish test",
			want: "This is a placeholder for the publish function",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := publish()
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}

// func TestKubedeploy(t *testing.T) {
// 	cases := []struct {
// 		name string
// 		want string
// 	}{
// 		{
// 			name: "Kube deploy test",
// 			want: "",
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			got := deploy().kubedir
// 			if !reflect.DeepEqual(got, c.want) {
// 				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
// 			}
// 		})
// 	}
// }
