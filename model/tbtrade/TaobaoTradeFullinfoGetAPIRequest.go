package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradefullinfogetAPIRequest 获取单笔交易的详细信息 API请求
// taobao.trade.fullinfo.get
//
// 获取单笔交易的详细信息
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包金额使用字段：tmall_coupon_fee（天猫商家订单使用的红包金额）
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaotradefullinfogetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// include_oaid
	_includeOaid string
	// 交易编号
	_tid int64
}

// NewTaobaotradefullinfogetRequest 初始化TaobaotradefullinfogetAPIRequest对象
func NewTaobaotradefullinfogetRequest() *TaobaotradefullinfogetAPIRequest {
	return &TaobaotradefullinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradefullinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.fullinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradefullinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradefullinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaotradefullinfogetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotradefullinfogetAPIRequest) GetFields() string {
	return r._fields
}

// SetIncludeOaid is IncludeOaid Setter
// include_oaid
func (r *TaobaotradefullinfogetAPIRequest) SetIncludeOaid(_includeOaid string) error {
	r._includeOaid = _includeOaid
	r.Set("include_oaid", _includeOaid)
	return nil
}

// GetIncludeOaid IncludeOaid Getter
func (r TaobaotradefullinfogetAPIRequest) GetIncludeOaid() string {
	return r._includeOaid
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaotradefullinfogetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradefullinfogetAPIRequest) GetTid() int64 {
	return r._tid
}
