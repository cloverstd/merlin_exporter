package merlin_client

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Info struct {
	Uptime      time.Time
	Temperature map[string]float64
	OSInfo      *OSInfo
}

func New(host, username, password string) (*MerlinClient, error) {
	jar, _ := cookiejar.New(nil)
	client := &MerlinClient{
		client: &http.Client{
			Jar: jar,
		},
		host:     host,
		username: username,
		password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Login(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

type MerlinClient struct {
	client   *http.Client
	username string
	password string
	host     string
}

func (mc *MerlinClient) Login(ctx context.Context) error {
	form := url.Values{}
	form.Set("login_authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", mc.username, mc.password))))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, mc.renderURL("login.cgi"), bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Referer", mc.renderURL(""))
	resp, err := mc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http call failed, %d", resp.StatusCode)
	}
	return nil
}

func (mc *MerlinClient) renderURL(path string) string {
	return fmt.Sprintf("http://%s/%s", mc.host, strings.TrimPrefix(path, "/"))
}

func (mc *MerlinClient) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := mc.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("http call failed with http code %d", resp.StatusCode)
	}
	return resp, nil
}

func (mc *MerlinClient) collect() (*Info, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := mc.Login(ctx); err != nil {
		return nil, err
	}
	var info Info
	uptime, err := mc.Uptime(ctx)
	if err != nil {
		return nil, err
	}
	info.Uptime = time.Now().Add(time.Duration(uptime) * time.Second)

	temperature, err := mc.Temperature(ctx)
	if err != nil {
		return nil, err
	}
	info.Temperature = temperature

	os, err := mc.OSInfo(ctx)
	if err != nil {
		return nil, err
	}
	info.OSInfo = os
	return &info, nil
}

func (mc *MerlinClient) Loop(interval time.Duration, fn func(info *Info)) {
	timer := time.NewTimer(1)
	defer timer.Stop()

	for range timer.C {
		collect, err := mc.collect()
		if err != nil {
			log.Printf("collect failed, %v\n", err)
		} else {
			fn(collect)
		}
		timer.Reset(interval)
	}
}
