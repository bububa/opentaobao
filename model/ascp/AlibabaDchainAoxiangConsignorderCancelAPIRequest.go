package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignordercancelAPIRequest 自动流转单据取消仓内发货 API请求
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
type AlibabadchainaoxiangconsignordercancelAPIRequest struct {
	model.Params
	// 取消仓内发货入参
	_cancelConsignorderRequest *CancelConsignOrderRequest
}

// NewAlibabadchainaoxiangconsignordercancelRequest 初始化AlibabadchainaoxiangconsignordercancelAPIRequest对象
func NewAlibabadchainaoxiangconsignordercancelRequest() *AlibabadchainaoxiangconsignordercancelAPIRequest {
	return &AlibabadchainaoxiangconsignordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangconsignordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangconsignordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangconsignordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelConsignorderRequest is CancelConsignorderRequest Setter
// 取消仓内发货入参
func (r *AlibabadchainaoxiangconsignordercancelAPIRequest) SetCancelConsignorderRequest(_cancelConsignorderRequest *CancelConsignOrderRequest) error {
	r._cancelConsignorderRequest = _cancelConsignorderRequest
	r.Set("cancel_consignorder_request", _cancelConsignorderRequest)
	return nil
}

// GetCancelConsignorderRequest CancelConsignorderRequest Getter
func (r AlibabadchainaoxiangconsignordercancelAPIRequest) GetCancelConsignorderRequest() *CancelConsignOrderRequest {
	return r._cancelConsignorderRequest
}
