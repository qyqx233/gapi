package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_executeCmd(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()
	t.Log(executeCmd(ctx, "dir", "D:\\download\\httpdownload"))
}

func Test(t *testing.T) {
	url := "http://localhost:3000/executeCmd"
	var rs ExecuteCmdRs
	rq := ExecuteCmdRq{
		Cmd:     "dir",
		Timeout: 1,
	}

	jsonStr, err := json.Marshal(rq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// defer req.Body.Close()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &rs)
	t.Log(rs.Err)
	t.Log(rs.Out)
	t.Log(rs.Msg)
}
