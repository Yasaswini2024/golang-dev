package main

import (
  "github.com/losandreas/scrapper"
  "github.com/yashaswini2024/notepad"
  "exec"
  "os"
  "fmt"
  "testing"
)

func TestOpen(t *testing.T) {
    err := Open()
    if err != nil {
        t.Errorf("Error: %v", err)
    }
}
