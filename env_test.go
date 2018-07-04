package nova

import "testing"

func TestEnv_IsProduction(t *testing.T) {
	var n Env
	if n.IsProduction() {
		t.Error("failed 1")
	}
	n = Env("production")
	if !n.IsProduction() {
		t.Error("failed 2")
	}
}

func TestEnv_IsDevelopment(t *testing.T) {
	var n Env
	if !n.IsDevelopment() {
		t.Error("failed 1")
	}
	n = Env("development")
	if n.IsProduction() {
		t.Error("failed 2")
	}
	if !n.IsDevelopment() {
		t.Error("failed 1")
	}
}

func TestEnv_IsTest(t *testing.T) {
	var n Env
	if n.IsTest() {
		t.Error("failed 1")
	}
	n = Env("test")
	if !n.IsTest() {
		t.Error("failed 2")
	}
}
