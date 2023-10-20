package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmhistoryouidgetAPIRequest 根据buyerNick获取ouid API请求
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
type TaobaocrmhistoryouidgetAPIRequest struct {
	model.Params
	// 买家淘宝Nick
	_buyerNick string
}

// NewTaobaocrmhistoryouidgetRequest 初始化TaobaocrmhistoryouidgetAPIRequest对象
func NewTaobaocrmhistoryouidgetRequest() *TaobaocrmhistoryouidgetAPIRequest {
	return &TaobaocrmhistoryouidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmhistoryouidgetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.history.ouid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmhistoryouidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmhistoryouidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家淘宝Nick
func (r *TaobaocrmhistoryouidgetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaocrmhistoryouidgetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
