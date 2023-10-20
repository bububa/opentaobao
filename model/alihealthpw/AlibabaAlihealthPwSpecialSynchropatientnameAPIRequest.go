package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwspecialsynchropatientnameAPIRequest 同步患者姓名至阿里健康 API请求
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
type AlibabaalihealthpwspecialsynchropatientnameAPIRequest struct {
	model.Params
	// 入参
	_body *SynchroPatientNameDto
}

// NewAlibabaalihealthpwspecialsynchropatientnameRequest 初始化AlibabaalihealthpwspecialsynchropatientnameAPIRequest对象
func NewAlibabaalihealthpwspecialsynchropatientnameRequest() *AlibabaalihealthpwspecialsynchropatientnameAPIRequest {
	return &AlibabaalihealthpwspecialsynchropatientnameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwspecialsynchropatientnameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchropatientname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwspecialsynchropatientnameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwspecialsynchropatientnameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaalihealthpwspecialsynchropatientnameAPIRequest) SetBody(_body *SynchroPatientNameDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwspecialsynchropatientnameAPIRequest) GetBody() *SynchroPatientNameDto {
	return r._body
}
