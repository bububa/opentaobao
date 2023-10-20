package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWtCifCoopOsstokenGetAPIRequest 获取oss签名接口 API请求
// alibaba.wt.cif.coop.osstoken.get
//
// 商家合作上传oss图片获取token接口
type AlibabaWtCifCoopOsstokenGetAPIRequest struct {
	model.Params
	// 调用方的应用名
	_appName string
	// 系统分配的source
	_source string
	// 系统分配的biz
	_biz string
}

// NewAlibabaWtCifCoopOsstokenGetRequest 初始化AlibabaWtCifCoopOsstokenGetAPIRequest对象
func NewAlibabaWtCifCoopOsstokenGetRequest() *AlibabaWtCifCoopOsstokenGetAPIRequest {
	return &AlibabaWtCifCoopOsstokenGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) Reset() {
	r._appName = ""
	r._source = ""
	r._biz = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wt.cif.coop.osstoken.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppName is AppName Setter
// 调用方的应用名
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetAppName() string {
	return r._appName
}

// SetSource is Source Setter
// 系统分配的source
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetSource() string {
	return r._source
}

// SetBiz is Biz Setter
// 系统分配的biz
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetBiz() string {
	return r._biz
}

var poolAlibabaWtCifCoopOsstokenGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWtCifCoopOsstokenGetRequest()
	},
}

// GetAlibabaWtCifCoopOsstokenGetRequest 从 sync.Pool 获取 AlibabaWtCifCoopOsstokenGetAPIRequest
func GetAlibabaWtCifCoopOsstokenGetAPIRequest() *AlibabaWtCifCoopOsstokenGetAPIRequest {
	return poolAlibabaWtCifCoopOsstokenGetAPIRequest.Get().(*AlibabaWtCifCoopOsstokenGetAPIRequest)
}

// ReleaseAlibabaWtCifCoopOsstokenGetAPIRequest 将 AlibabaWtCifCoopOsstokenGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWtCifCoopOsstokenGetAPIRequest(v *AlibabaWtCifCoopOsstokenGetAPIRequest) {
	v.Reset()
	poolAlibabaWtCifCoopOsstokenGetAPIRequest.Put(v)
}
