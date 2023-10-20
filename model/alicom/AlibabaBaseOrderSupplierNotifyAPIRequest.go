package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaseOrderSupplierNotifyAPIRequest 阿里通信运营商信息回传 API请求
// alibaba.base.order.supplier.notify
//
// 接收阿里通信流量运营商信息回传
type AlibabaBaseOrderSupplierNotifyAPIRequest struct {
	model.Params
	// 入参对象
	_paramFlowSuppllierNotifyModel *FlowSuppllierNotifyModel
}

// NewAlibabaBaseOrderSupplierNotifyRequest 初始化AlibabaBaseOrderSupplierNotifyAPIRequest对象
func NewAlibabaBaseOrderSupplierNotifyRequest() *AlibabaBaseOrderSupplierNotifyAPIRequest {
	return &AlibabaBaseOrderSupplierNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaseOrderSupplierNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.base.order.supplier.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaseOrderSupplierNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaseOrderSupplierNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFlowSuppllierNotifyModel is ParamFlowSuppllierNotifyModel Setter
// 入参对象
func (r *AlibabaBaseOrderSupplierNotifyAPIRequest) SetParamFlowSuppllierNotifyModel(_paramFlowSuppllierNotifyModel *FlowSuppllierNotifyModel) error {
	r._paramFlowSuppllierNotifyModel = _paramFlowSuppllierNotifyModel
	r.Set("param_flow_suppllier_notify_model", _paramFlowSuppllierNotifyModel)
	return nil
}

// GetParamFlowSuppllierNotifyModel ParamFlowSuppllierNotifyModel Getter
func (r AlibabaBaseOrderSupplierNotifyAPIRequest) GetParamFlowSuppllierNotifyModel() *FlowSuppllierNotifyModel {
	return r._paramFlowSuppllierNotifyModel
}
