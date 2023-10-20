package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniubuyertaggetAPIRequest 判断买家是否有某些标 API请求
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
type TaobaoqianniubuyertaggetAPIRequest struct {
	model.Params
	// 买家nick
	_buyerNick string
	// 支持的表，多个tag用英文逗号切割
	_tagList string
}

// NewTaobaoqianniubuyertaggetRequest 初始化TaobaoqianniubuyertaggetAPIRequest对象
func NewTaobaoqianniubuyertaggetRequest() *TaobaoqianniubuyertaggetAPIRequest {
	return &TaobaoqianniubuyertaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniubuyertaggetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.buyer.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniubuyertaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniubuyertaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 买家nick
func (r *TaobaoqianniubuyertaggetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoqianniubuyertaggetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetTagList is TagList Setter
// 支持的表，多个tag用英文逗号切割
func (r *TaobaoqianniubuyertaggetAPIRequest) SetTagList(_tagList string) error {
	r._tagList = _tagList
	r.Set("tag_list", _tagList)
	return nil
}

// GetTagList TagList Getter
func (r TaobaoqianniubuyertaggetAPIRequest) GetTagList() string {
	return r._tagList
}
