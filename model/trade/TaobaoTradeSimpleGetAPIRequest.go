package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeSimpleGetAPIRequest 获取交易订单的简易信息 API请求
// taobao.trade.simple.get
//
// 获取CRM单表交易的详情
// <br/>1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// <br/>2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// <br/>3. 获取红包优惠金额可以使用字段 coupon_fee
// <br/>4. 请按需获取字段，减少TOP系统的压力
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
type TaobaoTradeSimpleGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// NewTaobaoTradeSimpleGetRequest 初始化TaobaoTradeSimpleGetAPIRequest对象
func NewTaobaoTradeSimpleGetRequest() *TaobaoTradeSimpleGetAPIRequest {
	return &TaobaoTradeSimpleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeSimpleGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.simple.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeSimpleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoTradeSimpleGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTradeSimpleGetAPIRequest) GetFields() string {
	return r._fields
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaoTradeSimpleGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeSimpleGetAPIRequest) GetTid() int64 {
	return r._tid
}
