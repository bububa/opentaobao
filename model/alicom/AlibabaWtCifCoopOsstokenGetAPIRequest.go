package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawtcifcooposstokengetAPIRequest 获取oss签名接口 API请求
// alibaba.wt.cif.coop.osstoken.get
//
// 商家合作上传oss图片获取token接口
type AlibabawtcifcooposstokengetAPIRequest struct {
	model.Params
	// 调用方的应用名
	_appName string
	// 系统分配的source
	_source string
	// 系统分配的biz
	_biz string
}

// NewAlibabawtcifcooposstokengetRequest 初始化AlibabawtcifcooposstokengetAPIRequest对象
func NewAlibabawtcifcooposstokengetRequest() *AlibabawtcifcooposstokengetAPIRequest {
	return &AlibabawtcifcooposstokengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawtcifcooposstokengetAPIRequest) GetApiMethodName() string {
	return "alibaba.wt.cif.coop.osstoken.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawtcifcooposstokengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawtcifcooposstokengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppName is AppName Setter
// 调用方的应用名
func (r *AlibabawtcifcooposstokengetAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibabawtcifcooposstokengetAPIRequest) GetAppName() string {
	return r._appName
}

// SetSource is Source Setter
// 系统分配的source
func (r *AlibabawtcifcooposstokengetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabawtcifcooposstokengetAPIRequest) GetSource() string {
	return r._source
}

// SetBiz is Biz Setter
// 系统分配的biz
func (r *AlibabawtcifcooposstokengetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r AlibabawtcifcooposstokengetAPIRequest) GetBiz() string {
	return r._biz
}
