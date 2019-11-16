package request

import (
	"fmt"
	"net/http"
)

func (rq *Request) SetDebug(debug bool) *Request {
	rq.d = debug
	return rq
}

func (rq *Request) SetBaseUrl(baseUrl string) *Request {
	rq.b = baseUrl
	return rq
}

func (rq *Request) SetUrl(apiUrl string) *Request {
	rq.u = apiUrl
	return rq
}

func (rq *Request) SetUrlf(urlFormat string, v ...interface{}) *Request {
	rq.u = fmt.Sprintf(urlFormat, v...)
	return rq
}

func (rq *Request) SetMethod(method string) *Request {
	rq.m = method
	return rq
}
// 设置超时时间,超时包括连接时间，任何重定向和读取响应正文,0表示不会超时
func (rq *Request) SetTimeOut(second int) *Request {
	rq.p.TimeOut = second
	return rq
}

func (rq *Request) SetProxy(proxy string) *Request {
	rq.p.Proxy = proxy
	return rq
}
// map | string(query | lines | json)
func (rq *Request) SetParam(v interface{}) *Request {
	rq.p.Param = toMSS(v)
	return rq
}

func (rq *Request) AddParam(k string, v interface{}) *Request {
	rq.p.Param[k] = anyToString(v)
	return rq
}
// map | string(query | lines | json)
func (rq *Request) SetHeader(v interface{}) *Request {
	rq.p.Header = toMSS(v)
	return rq
}

func (rq *Request) AddHeader(k string, v interface{}) *Request {
	rq.p.Header[k] = anyToString(v)
	return rq
}

func (rq *Request) AddCookies(cookies ...*Cookie) *Request {
	if cookies != nil{
		rq.p.Cookie = append(rq.p.Cookie,cookies...)
	}
	return rq
}

func (rq *Request) AddCookieSimple(k string, v interface{}) *Request {
	rq.p.Cookie = append(rq.p.Cookie,&Cookie{
		Name: k,
		Value: anyToString(v),
	})
	return rq
}

func (rq *Request) SetData(v interface{}) *Request {
	rq.p.Data = toMSI(v)
	return rq
}

func (rq *Request) AddData(k string, v interface{}) *Request {
	rq.p.Data[k] = v
	return rq
}

func (rq *Request) SetFormData(v interface{}) *Request {
	rq.p.FormData = toMSS(v)
	return rq
}

func (rq *Request) AddFormData(k string, v interface{}) *Request {
	rq.p.FormData[k] = anyToString(v)
	return rq
}
// todo
func (rq *Request) SetFiles(filePaths ...string) *Request {
	rq.p.Files = filePaths
	return rq
}
// todo
func (rq *Request) AddFile(filePath string) *Request {
	rq.p.Files = append(rq.p.Files, filePath)
	return rq
}
// 设置 Transport
func (rq *Request) SetTransport(transport *http.Transport) *Request {
	rq.t = transport
	return rq
}

func (rq *Request) SetRequest(r *http.Request) *Request {
	rq.r = r
	return rq
}

func (rq *Request) SetClient(c *http.Client) *Request {
	rq.c = c
	return rq
}

func (rq *Request) SetOption(option *Option) *Request {
	rq.p = option
	return rq
}


func (rq *Request) DoHttp() IResponse {
	return rq.send()
}

func (rq *Request) DoGet(apiUrl string, param interface{}) IResponse {
	rq.SetMethod(GET)
	rq.AddHeader(CONTENT_TYPE,APPLICATION_JSON)
	rq.SetUrl(apiUrl)
	rq.SetParam(param)
	return rq.send()
}

func (rq *Request) DoPost(apiUrl string, data interface{}) IResponse {
	rq.SetMethod(POST)
	rq.AddHeader(CONTENT_TYPE,APPLICATION_JSON)
	rq.SetUrl(apiUrl)
	rq.SetData(data)
	return rq.send()
}

func (rq *Request) DoPostForm(apiUrl string, formData interface{}) IResponse {
	rq.SetMethod(POST)
	rq.AddHeader(CONTENT_TYPE,X_WWW_FORM_URLENCODED)
	rq.SetUrl(apiUrl)
	rq.SetFormData(formData)
	return rq.send()
}

