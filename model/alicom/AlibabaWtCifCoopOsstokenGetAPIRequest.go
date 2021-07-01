package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWtCifCoopOsstokenGetAPIRequest
获取oss签名接口 API请求
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wt.cif.coop.osstoken.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppName Setter
// 调用方的应用名
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// Get AppName Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetAppName() string {
	return r._appName
}

// Set is Source Setter
// 系统分配的source
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetSource() string {
	return r._source
}

// Set is Biz Setter
// 系统分配的biz
func (r *AlibabaWtCifCoopOsstokenGetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// Get Biz Getter
func (r AlibabaWtCifCoopOsstokenGetAPIRequest) GetBiz() string {
	return r._biz
}
