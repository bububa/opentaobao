package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgrademktmemberqueryAPIRequest 会员等级营销-会员关系查询 API请求
// taobao.crm.grademkt.member.query
//
// 商家通过该接口查询线上店铺会员。
type TaobaocrmgrademktmemberqueryAPIRequest struct {
	model.Params
	// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 系统属性，json格式
	_feather string
	// 会员nick
	_buyerNick string
}

// NewTaobaocrmgrademktmemberqueryRequest 初始化TaobaocrmgrademktmemberqueryAPIRequest对象
func NewTaobaocrmgrademktmemberqueryRequest() *TaobaocrmgrademktmemberqueryAPIRequest {
	return &TaobaocrmgrademktmemberqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaocrmgrademktmemberqueryAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 系统属性，json格式
func (r *TaobaocrmgrademktmemberqueryAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetFeather() string {
	return r._feather
}

// SetBuyerNick is BuyerNick Setter
// 会员nick
func (r *TaobaocrmgrademktmemberqueryAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaocrmgrademktmemberqueryAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
