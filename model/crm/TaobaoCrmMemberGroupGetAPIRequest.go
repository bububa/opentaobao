package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmembergroupgetAPIRequest 获取买家身上的标签 API请求
// taobao.crm.member.group.get
//
// 获取买家身上的标签，不返回标签的总人数
type TaobaocrmmembergroupgetAPIRequest struct {
	model.Params
	// 会员Nick
	_buyerNick string
}

// NewTaobaocrmmembergroupgetRequest 初始化TaobaocrmmembergroupgetAPIRequest对象
func NewTaobaocrmmembergroupgetRequest() *TaobaocrmmembergroupgetAPIRequest {
	return &TaobaocrmmembergroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmmembergroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.member.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmmembergroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmmembergroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 会员Nick
func (r *TaobaocrmmembergroupgetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaocrmmembergroupgetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
