package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOuidGetAPIRequest 根据buyerNick获取ouid API请求
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOuidGetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaoCrmHistoryOuidGetRequest 初始化TaobaoCrmHistoryOuidGetAPIRequest对象
func NewTaobaoCrmHistoryOuidGetRequest() *TaobaoCrmHistoryOuidGetAPIRequest {
	return &TaobaoCrmHistoryOuidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.ouid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaoCrmHistoryOuidGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmHistoryOuidGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
