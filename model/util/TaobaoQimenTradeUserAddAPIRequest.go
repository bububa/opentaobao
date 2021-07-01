package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTradeUserAddAPIRequest
添加奇门订单链路用户 API请求
taobao.qimen.trade.user.add

添加奇门订单链路用户 */
type TaobaoQimenTradeUserAddAPIRequest struct {
	model.Params
	// 商家备注
	_memo string
}

// NewTaobaoQimenTradeUserAddRequest 初始化TaobaoQimenTradeUserAddAPIRequest对象
func NewTaobaoQimenTradeUserAddRequest() *TaobaoQimenTradeUserAddAPIRequest {
	return &TaobaoQimenTradeUserAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUserAddAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUserAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Memo Setter
// 商家备注
func (r *TaobaoQimenTradeUserAddAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// Get Memo Getter
func (r TaobaoQimenTradeUserAddAPIRequest) GetMemo() string {
	return r._memo
}
