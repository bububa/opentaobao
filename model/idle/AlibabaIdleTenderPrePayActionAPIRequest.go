package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderprepayactionAPIRequest 服务商预付款完成接口 API请求
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
type AlibabaidletenderprepayactionAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *TenderPrePayCmd
}

// NewAlibabaidletenderprepayactionRequest 初始化AlibabaidletenderprepayactionAPIRequest对象
func NewAlibabaidletenderprepayactionRequest() *AlibabaidletenderprepayactionAPIRequest {
	return &AlibabaidletenderprepayactionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderprepayactionAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.pre.pay.action"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderprepayactionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderprepayactionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaidletenderprepayactionAPIRequest) SetParam0(_param0 *TenderPrePayCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidletenderprepayactionAPIRequest) GetParam0() *TenderPrePayCmd {
	return r._param0
}
