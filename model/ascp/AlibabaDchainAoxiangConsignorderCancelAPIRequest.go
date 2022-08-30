package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderCancelAPIRequest 自动流转单据取消仓内发货 API请求
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
type AlibabaDchainAoxiangConsignorderCancelAPIRequest struct {
	model.Params
	// 取消仓内发货入参
	_cancelConsignorderRequest *CancelConsignOrderRequest
}

// NewAlibabaDchainAoxiangConsignorderCancelRequest 初始化AlibabaDchainAoxiangConsignorderCancelAPIRequest对象
func NewAlibabaDchainAoxiangConsignorderCancelRequest() *AlibabaDchainAoxiangConsignorderCancelAPIRequest {
	return &AlibabaDchainAoxiangConsignorderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCancelConsignorderRequest is CancelConsignorderRequest Setter
// 取消仓内发货入参
func (r *AlibabaDchainAoxiangConsignorderCancelAPIRequest) SetCancelConsignorderRequest(_cancelConsignorderRequest *CancelConsignOrderRequest) error {
	r._cancelConsignorderRequest = _cancelConsignorderRequest
	r.Set("cancel_consignorder_request", _cancelConsignorderRequest)
	return nil
}

// GetCancelConsignorderRequest CancelConsignorderRequest Getter
func (r AlibabaDchainAoxiangConsignorderCancelAPIRequest) GetCancelConsignorderRequest() *CancelConsignOrderRequest {
	return r._cancelConsignorderRequest
}
