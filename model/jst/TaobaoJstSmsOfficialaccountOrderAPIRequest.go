package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountOrderAPIRequest
聚石塔订购公众号 API请求
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口 */
type TaobaoJstSmsOfficialaccountOrderAPIRequest struct {
	model.Params
	// 聚石塔公众号订购
	_orderOfficialAccountRequest *OrderOfficialAccountRequest
}

// NewTaobaoJstSmsOfficialaccountOrderRequest 初始化TaobaoJstSmsOfficialaccountOrderAPIRequest对象
func NewTaobaoJstSmsOfficialaccountOrderRequest() *TaobaoJstSmsOfficialaccountOrderAPIRequest {
	return &TaobaoJstSmsOfficialaccountOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOrderAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.officialaccount.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOrderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderOfficialAccountRequest Setter
// 聚石塔公众号订购
func (r *TaobaoJstSmsOfficialaccountOrderAPIRequest) SetOrderOfficialAccountRequest(_orderOfficialAccountRequest *OrderOfficialAccountRequest) error {
	r._orderOfficialAccountRequest = _orderOfficialAccountRequest
	r.Set("order_official_account_request", _orderOfficialAccountRequest)
	return nil
}

// Get OrderOfficialAccountRequest Getter
func (r TaobaoJstSmsOfficialaccountOrderAPIRequest) GetOrderOfficialAccountRequest() *OrderOfficialAccountRequest {
	return r._orderOfficialAccountRequest
}
