package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmhistoryomidgetAPIRequest 根据buyerNick获取omid API请求
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
type TaobaocrmhistoryomidgetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaocrmhistoryomidgetRequest 初始化TaobaocrmhistoryomidgetAPIRequest对象
func NewTaobaocrmhistoryomidgetRequest() *TaobaocrmhistoryomidgetAPIRequest {
	return &TaobaocrmhistoryomidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmhistoryomidgetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.omid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmhistoryomidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmhistoryomidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaocrmhistoryomidgetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaocrmhistoryomidgetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
