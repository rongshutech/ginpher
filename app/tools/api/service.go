package api

import (
	"fmt"
	"strings"

	"github.com/vgmdj/utils/httplib"
	"github.com/vgmdj/utils/logger"
)

const (
	PINGOK = "OK"
)

type Config struct {
	BaseURL string
}

type Service struct {
	baseURL string
	cli     *httplib.Client
}

func NewService(c *Config) (srv *Service) {
	return &Service{
		baseURL: strings.TrimRight(c.BaseURL, "/"),
		cli:     httplib.UniqueClient(httplib.DefaultClientConf),
	}
}

func (srv *Service) PING(path string) error {
	path = "/" + strings.TrimLeft(path, "/")
	headers := make(map[string]string, 0)
	headers[httplib.ResponseResultContentType] = httplib.MIMEPlain
	resp := ""
	err := srv.cli.Get(srv.baseURL+path, nil, resp, headers)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if resp != PINGOK {
		return fmt.Errorf("service can not reach")
	}

	return nil
}

func (srv *Service) GET(path string, query map[string]interface{}, resp interface{}, headers map[string]string) error {
	path = "/" + strings.TrimLeft(path, "/")
	if len(headers) == 0 {
		headers = make(map[string]string)
	}
	headers[httplib.ResponseResultContentType] = httplib.MIMEJSON

	return srv.cli.Get(srv.baseURL+path, query, resp, headers)
}

func (srv *Service) POST(path string, body interface{}, resp interface{}, headers map[string]string) error {
	path = "/" + strings.TrimLeft(path, "/")
	if len(headers) == 0 {
		headers = make(map[string]string)
	}
	headers[httplib.ResponseResultContentType] = httplib.MIMEJSON

	return srv.cli.PostJSON(srv.baseURL+path, body, resp, headers)
}

func (srv *Service) POSTFORM(path string, formValues map[string]string, resp interface{}, headers map[string]string) error {
	path = "/" + strings.TrimLeft(path, "/")
	if len(headers) == 0 {
		headers = make(map[string]string)
	}
	headers[httplib.ResponseResultContentType] = httplib.MIMEJSON

	return srv.cli.PostForm(srv.baseURL+path, formValues, resp, headers)
}
