package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqUrlScanAPIRequest 恶意网址检测接口 API请求
// alibaba.security.jaq.url.scan
//
// url扫描接口
type AlibabaSecurityJaqUrlScanAPIRequest struct {
	model.Params
	// 扫描参数
	_paramUrlScanParamList *UrlScanParamList
}

// NewAlibabaSecurityJaqUrlScanRequest 初始化AlibabaSecurityJaqUrlScanAPIRequest对象
func NewAlibabaSecurityJaqUrlScanRequest() *AlibabaSecurityJaqUrlScanAPIRequest {
	return &AlibabaSecurityJaqUrlScanAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqUrlScanAPIRequest) Reset() {
	r._paramUrlScanParamList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqUrlScanAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.url.scan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqUrlScanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqUrlScanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUrlScanParamList is ParamUrlScanParamList Setter
// 扫描参数
func (r *AlibabaSecurityJaqUrlScanAPIRequest) SetParamUrlScanParamList(_paramUrlScanParamList *UrlScanParamList) error {
	r._paramUrlScanParamList = _paramUrlScanParamList
	r.Set("param_url_scan_param_list", _paramUrlScanParamList)
	return nil
}

// GetParamUrlScanParamList ParamUrlScanParamList Getter
func (r AlibabaSecurityJaqUrlScanAPIRequest) GetParamUrlScanParamList() *UrlScanParamList {
	return r._paramUrlScanParamList
}

var poolAlibabaSecurityJaqUrlScanAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqUrlScanRequest()
	},
}

// GetAlibabaSecurityJaqUrlScanRequest 从 sync.Pool 获取 AlibabaSecurityJaqUrlScanAPIRequest
func GetAlibabaSecurityJaqUrlScanAPIRequest() *AlibabaSecurityJaqUrlScanAPIRequest {
	return poolAlibabaSecurityJaqUrlScanAPIRequest.Get().(*AlibabaSecurityJaqUrlScanAPIRequest)
}

// ReleaseAlibabaSecurityJaqUrlScanAPIRequest 将 AlibabaSecurityJaqUrlScanAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqUrlScanAPIRequest(v *AlibabaSecurityJaqUrlScanAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqUrlScanAPIRequest.Put(v)
}
