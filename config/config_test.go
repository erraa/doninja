package config

import (
	"fmt"
	"testing"
)

func TestConfigt(test *testing.T) {
	ReadConfig()
	fmt.Println(Config)
}
