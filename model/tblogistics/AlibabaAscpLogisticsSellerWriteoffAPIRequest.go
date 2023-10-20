package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellerwriteoffAPIRequest 商家配送核销 API请求
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
type AlibabaascplogisticssellerwriteoffAPIRequest struct {
	model.Params
	// 核销码
	_receiveCode string
	// 所要核销的核销物流单ID，在alibaba.ascp.logistics.seller.orders.get中获取。
	_lpOrderId int64
}

// NewAlibabaascplogisticssellerwriteoffRequest 初始化AlibabaascplogisticssellerwriteoffAPIRequest对象
func NewAlibabaascplogisticssellerwriteoffRequest() *AlibabaascplogisticssellerwriteoffAPIRequest {
	return &AlibabaascplogisticssellerwriteoffAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticssellerwriteoffAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.writeoff"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticssellerwriteoffAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticssellerwriteoffAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiveCode is ReceiveCode Setter
// 核销码
func (r *AlibabaascplogisticssellerwriteoffAPIRequest) SetReceiveCode(_receiveCode string) error {
	r._receiveCode = _receiveCode
	r.Set("receive_code", _receiveCode)
	return nil
}

// GetReceiveCode ReceiveCode Getter
func (r AlibabaascplogisticssellerwriteoffAPIRequest) GetReceiveCode() string {
	return r._receiveCode
}

// SetLpOrderId is LpOrderId Setter
// 所要核销的核销物流单ID，在alibaba.ascp.logistics.seller.orders.get中获取。
func (r *AlibabaascplogisticssellerwriteoffAPIRequest) SetLpOrderId(_lpOrderId int64) error {
	r._lpOrderId = _lpOrderId
	r.Set("lp_order_id", _lpOrderId)
	return nil
}

// GetLpOrderId LpOrderId Getter
func (r AlibabaascplogisticssellerwriteoffAPIRequest) GetLpOrderId() int64 {
	return r._lpOrderId
}
