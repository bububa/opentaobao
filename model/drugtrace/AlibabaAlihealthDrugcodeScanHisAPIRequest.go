package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodescanhisAPIRequest 企业查询扫码历史 API请求
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
type AlibabaalihealthdrugcodescanhisAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaalihealthdrugcodescanhisRequest 初始化AlibabaalihealthdrugcodescanhisAPIRequest对象
func NewAlibabaalihealthdrugcodescanhisRequest() *AlibabaalihealthdrugcodescanhisAPIRequest {
	return &AlibabaalihealthdrugcodescanhisAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodescanhisAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.scan.his"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodescanhisAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodescanhisAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugcodescanhisAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodescanhisAPIRequest) GetCode() string {
	return r._code
}
