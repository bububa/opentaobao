package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangReceiverinfoQueryAPIRequest 供应链优仓即时配隐私小号查询 API请求
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
type AlibabaDchainAoxiangReceiverinfoQueryAPIRequest struct {
	model.Params
	// 收件人查询条件
	_orderPrivacyReceiverQuery *OrderPrivacyReceiverQuery
}

// NewAlibabaDchainAoxiangReceiverinfoQueryRequest 初始化AlibabaDchainAoxiangReceiverinfoQueryAPIRequest对象
func NewAlibabaDchainAoxiangReceiverinfoQueryRequest() *AlibabaDchainAoxiangReceiverinfoQueryAPIRequest {
	return &AlibabaDchainAoxiangReceiverinfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangReceiverinfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.receiverinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangReceiverinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangReceiverinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderPrivacyReceiverQuery is OrderPrivacyReceiverQuery Setter
// 收件人查询条件
func (r *AlibabaDchainAoxiangReceiverinfoQueryAPIRequest) SetOrderPrivacyReceiverQuery(_orderPrivacyReceiverQuery *OrderPrivacyReceiverQuery) error {
	r._orderPrivacyReceiverQuery = _orderPrivacyReceiverQuery
	r.Set("order_privacy_receiver_query", _orderPrivacyReceiverQuery)
	return nil
}

// GetOrderPrivacyReceiverQuery OrderPrivacyReceiverQuery Getter
func (r AlibabaDchainAoxiangReceiverinfoQueryAPIRequest) GetOrderPrivacyReceiverQuery() *OrderPrivacyReceiverQuery {
	return r._orderPrivacyReceiverQuery
}
