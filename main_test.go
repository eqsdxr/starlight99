package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_generateNewExpression(t *testing.T) {
	totalScores := []int{1, 50, 150, 500, 5000}
	for _, v := range totalScores {
		exp := generateNextExpression(v)
		if exp.First < 1 || exp.Second < 1 || exp.Damage < 1 || exp.Result < 1 {
			t.Error("Some field from result is less than 1")
		}
	}
}

func Test_saveConfig(t *testing.T) {
	cfg := Config{99, 1, 0}
	err := saveConfig(&cfg, configPath)
	if err != nil {
		t.Error("saveConfig error")
	}
}

func Test_readConfig(t *testing.T) {
	cfg, err := readConfig(configPath)
	if err != nil {
		t.Error("readConfig error")
	}
	if cfg == nil {
		t.Error("cfg returned by readConfig is nil")
	}
}

func Test_getInput(t *testing.T) {
	input := "1\n"
	expected := 1

	scanner := *bufio.NewScanner(strings.NewReader(input))

	actual := getInput(&scanner)
	if actual != expected {
		t.Error("getInput error")
	}
}

func Test_setPlayerHealth(t *testing.T) {
	totalScores := []int{1, 50, 150, 500, 5000}
	for _, v := range totalScores {
		if health := setPlayerHealth(v); health < 500 {
			t.Error("Some field from result is less than 500")
		}
	}
}

func Test_getNextMonster(t *testing.T) {
	totalScores := []int{1, 50, 150, 500, 5000}
	for _, v := range totalScores {
		if m := getNextMonster(v); len(m.ASCII) < 50 || len(m.Name) < 2 ||
		m.DamageMax < 1 || m.DamageMin < 1 || m.HPMax < 1 || m.HPMin < 1 || m.HP < 1{
			t.Error("Some field from monster is incorrect")
		}
	}
}
