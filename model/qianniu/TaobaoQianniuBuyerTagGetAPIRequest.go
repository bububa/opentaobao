package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuBuyerTagGetAPIRequest 判断买家是否有某些标 API请求
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
type TaobaoQianniuBuyerTagGetAPIRequest struct {
	model.Params
	// 买家nick
	_buyerNick string
	// 支持的表，多个tag用英文逗号切割
	_tagList string
}

// NewTaobaoQianniuBuyerTagGetRequest 初始化TaobaoQianniuBuyerTagGetAPIRequest对象
func NewTaobaoQianniuBuyerTagGetRequest() *TaobaoQianniuBuyerTagGetAPIRequest {
	return &TaobaoQianniuBuyerTagGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuBuyerTagGetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.buyer.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuBuyerTagGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 买家nick
func (r *TaobaoQianniuBuyerTagGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoQianniuBuyerTagGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetTagList is TagList Setter
// 支持的表，多个tag用英文逗号切割
func (r *TaobaoQianniuBuyerTagGetAPIRequest) SetTagList(_tagList string) error {
	r._tagList = _tagList
	r.Set("tag_list", _tagList)
	return nil
}

// GetTagList TagList Getter
func (r TaobaoQianniuBuyerTagGetAPIRequest) GetTagList() string {
	return r._tagList
}
