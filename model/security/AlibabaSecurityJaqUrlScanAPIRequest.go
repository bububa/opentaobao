package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqurlscanAPIRequest 恶意网址检测接口 API请求
// alibaba.security.jaq.url.scan
//
// url扫描接口
type AlibabasecurityjaqurlscanAPIRequest struct {
	model.Params
	// 扫描参数
	_paramUrlScanParamList *UrlScanParamList
}

// NewAlibabasecurityjaqurlscanRequest 初始化AlibabasecurityjaqurlscanAPIRequest对象
func NewAlibabasecurityjaqurlscanRequest() *AlibabasecurityjaqurlscanAPIRequest {
	return &AlibabasecurityjaqurlscanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqurlscanAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.url.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqurlscanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqurlscanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUrlScanParamList is ParamUrlScanParamList Setter
// 扫描参数
func (r *AlibabasecurityjaqurlscanAPIRequest) SetParamUrlScanParamList(_paramUrlScanParamList *UrlScanParamList) error {
	r._paramUrlScanParamList = _paramUrlScanParamList
	r.Set("param_url_scan_param_list", _paramUrlScanParamList)
	return nil
}

// GetParamUrlScanParamList ParamUrlScanParamList Getter
func (r AlibabasecurityjaqurlscanAPIRequest) GetParamUrlScanParamList() *UrlScanParamList {
	return r._paramUrlScanParamList
}
