package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwspecialsynchronodeAPIRequest 合作方同步状态至阿里健康 API请求
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
type AlibabaalihealthpwspecialsynchronodeAPIRequest struct {
	model.Params
	// 状态信息入参
	_body *SnodeDto
}

// NewAlibabaalihealthpwspecialsynchronodeRequest 初始化AlibabaalihealthpwspecialsynchronodeAPIRequest对象
func NewAlibabaalihealthpwspecialsynchronodeRequest() *AlibabaalihealthpwspecialsynchronodeAPIRequest {
	return &AlibabaalihealthpwspecialsynchronodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwspecialsynchronodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchronode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwspecialsynchronodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwspecialsynchronodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 状态信息入参
func (r *AlibabaalihealthpwspecialsynchronodeAPIRequest) SetBody(_body *SnodeDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwspecialsynchronodeAPIRequest) GetBody() *SnodeDto {
	return r._body
}
