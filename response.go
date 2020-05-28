package http

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Response struct {
	err  error
	resp *http.Response
	body []byte
}

func (r *Response) Err() error {
	return r.err
}

func (r *Response) Response() *http.Response {
	return r.resp
}

func (r *Response) Val() []byte {
	if r.err == nil && r.body == nil {
		r.body, r.err = ioutil.ReadAll(r.resp.Body)
		defer r.resp.Body.Close()
	}
	return r.body
}

func (r *Response) String() string {
	return string(r.Val())
}

func (r *Response) Result() ([]byte, error) {
	return r.Val(), r.err
}

func (r *Response) JsonUnmarshal(v interface{}) error {
	data := r.Val()
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal(data, v)
}

func (r *Response) XmlUnmarshal(v interface{}) error {
	data := r.Val()
	if r.err != nil {
		return r.err
	}
	return xml.Unmarshal(data, v)
}
