package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmHistoryOmidGetAPIRequest 根据buyerNick获取omid API请求
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
type TaobaoCrmHistoryOmidGetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaoCrmHistoryOmidGetRequest 初始化TaobaoCrmHistoryOmidGetAPIRequest对象
func NewTaobaoCrmHistoryOmidGetRequest() *TaobaoCrmHistoryOmidGetAPIRequest {
	return &TaobaoCrmHistoryOmidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.omid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaoCrmHistoryOmidGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmHistoryOmidGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
