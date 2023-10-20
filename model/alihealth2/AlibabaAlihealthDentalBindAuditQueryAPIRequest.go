package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalbindauditqueryAPIRequest ISV查询绑定审核状态 API请求
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
type AlibabaalihealthdentalbindauditqueryAPIRequest struct {
	model.Params
	// 绑定ID列表
	_bindIds []string
}

// NewAlibabaalihealthdentalbindauditqueryRequest 初始化AlibabaalihealthdentalbindauditqueryAPIRequest对象
func NewAlibabaalihealthdentalbindauditqueryRequest() *AlibabaalihealthdentalbindauditqueryAPIRequest {
	return &AlibabaalihealthdentalbindauditqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalbindauditqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.bind.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalbindauditqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalbindauditqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindIds is BindIds Setter
// 绑定ID列表
func (r *AlibabaalihealthdentalbindauditqueryAPIRequest) SetBindIds(_bindIds []string) error {
	r._bindIds = _bindIds
	r.Set("bind_ids", _bindIds)
	return nil
}

// GetBindIds BindIds Getter
func (r AlibabaalihealthdentalbindauditqueryAPIRequest) GetBindIds() []string {
	return r._bindIds
}
