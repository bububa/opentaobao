package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangreceiverinfoqueryAPIRequest 供应链优仓即时配隐私小号查询 API请求
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
type AlibabadchainaoxiangreceiverinfoqueryAPIRequest struct {
	model.Params
	// 收件人查询条件
	_orderPrivacyReceiverQuery *OrderPrivacyReceiverQuery
}

// NewAlibabadchainaoxiangreceiverinfoqueryRequest 初始化AlibabadchainaoxiangreceiverinfoqueryAPIRequest对象
func NewAlibabadchainaoxiangreceiverinfoqueryRequest() *AlibabadchainaoxiangreceiverinfoqueryAPIRequest {
	return &AlibabadchainaoxiangreceiverinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangreceiverinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.receiverinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangreceiverinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangreceiverinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderPrivacyReceiverQuery is OrderPrivacyReceiverQuery Setter
// 收件人查询条件
func (r *AlibabadchainaoxiangreceiverinfoqueryAPIRequest) SetOrderPrivacyReceiverQuery(_orderPrivacyReceiverQuery *OrderPrivacyReceiverQuery) error {
	r._orderPrivacyReceiverQuery = _orderPrivacyReceiverQuery
	r.Set("order_privacy_receiver_query", _orderPrivacyReceiverQuery)
	return nil
}

// GetOrderPrivacyReceiverQuery OrderPrivacyReceiverQuery Getter
func (r AlibabadchainaoxiangreceiverinfoqueryAPIRequest) GetOrderPrivacyReceiverQuery() *OrderPrivacyReceiverQuery {
	return r._orderPrivacyReceiverQuery
}
