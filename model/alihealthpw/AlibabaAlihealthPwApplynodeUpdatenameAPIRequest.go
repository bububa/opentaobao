package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwapplynodeupdatenameAPIRequest 回调变更患者姓名 API请求
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
type AlibabaalihealthpwapplynodeupdatenameAPIRequest struct {
	model.Params
	// 回调入参
	_body *ModifyNameRo
}

// NewAlibabaalihealthpwapplynodeupdatenameRequest 初始化AlibabaalihealthpwapplynodeupdatenameAPIRequest对象
func NewAlibabaalihealthpwapplynodeupdatenameRequest() *AlibabaalihealthpwapplynodeupdatenameAPIRequest {
	return &AlibabaalihealthpwapplynodeupdatenameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwapplynodeupdatenameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.updatename"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwapplynodeupdatenameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwapplynodeupdatenameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 回调入参
func (r *AlibabaalihealthpwapplynodeupdatenameAPIRequest) SetBody(_body *ModifyNameRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwapplynodeupdatenameAPIRequest) GetBody() *ModifyNameRo {
	return r._body
}
