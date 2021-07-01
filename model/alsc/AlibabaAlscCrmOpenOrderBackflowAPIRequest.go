package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenOrderBackflowAPIRequest
订单回流接口 API请求
alibaba.alsc.crm.open.order.backflow

回流isv订单接口 */
type AlibabaAlscCrmOpenOrderBackflowAPIRequest struct {
	model.Params
	// 入参
	_paramOrderBackflowOpenReq *OrderBackflowOpenReq
}

// NewAlibabaAlscCrmOpenOrderBackflowRequest 初始化AlibabaAlscCrmOpenOrderBackflowAPIRequest对象
func NewAlibabaAlscCrmOpenOrderBackflowRequest() *AlibabaAlscCrmOpenOrderBackflowAPIRequest {
	return &AlibabaAlscCrmOpenOrderBackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOrderBackflowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenOrderBackflowAPIRequest) SetParamOrderBackflowOpenReq(_paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
	r._paramOrderBackflowOpenReq = _paramOrderBackflowOpenReq
	r.Set("param_order_backflow_open_req", _paramOrderBackflowOpenReq)
	return nil
}

// Get ParamOrderBackflowOpenReq Getter
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
	return r._paramOrderBackflowOpenReq
}
