package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeFullinfoGetAPIRequest 获取单笔交易的详细信息 API请求
// taobao.trade.fullinfo.get
//
// 获取单笔交易的详细信息
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包优惠金额可以使用字段 coupon_fee
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoTradeFullinfoGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// include_oaid
	_includeOaid string
	// 交易编号
	_tid int64
}

// NewTaobaoTradeFullinfoGetRequest 初始化TaobaoTradeFullinfoGetAPIRequest对象
func NewTaobaoTradeFullinfoGetRequest() *TaobaoTradeFullinfoGetAPIRequest {
	return &TaobaoTradeFullinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeFullinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.fullinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeFullinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoTradeFullinfoGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTradeFullinfoGetAPIRequest) GetFields() string {
	return r._fields
}

// SetIncludeOaid is IncludeOaid Setter
// include_oaid
func (r *TaobaoTradeFullinfoGetAPIRequest) SetIncludeOaid(_includeOaid string) error {
	r._includeOaid = _includeOaid
	r.Set("include_oaid", _includeOaid)
	return nil
}

// GetIncludeOaid IncludeOaid Getter
func (r TaobaoTradeFullinfoGetAPIRequest) GetIncludeOaid() string {
	return r._includeOaid
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaoTradeFullinfoGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeFullinfoGetAPIRequest) GetTid() int64 {
	return r._tid
}
