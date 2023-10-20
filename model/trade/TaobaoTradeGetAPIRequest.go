package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradegetAPIRequest 获取单笔交易的部分信息(性能高) API请求
// taobao.trade.get
//
// 获取单笔交易的部分信息
// &lt;br/&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;
type TaobaotradegetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// NewTaobaotradegetRequest 初始化TaobaotradegetAPIRequest对象
func NewTaobaotradegetRequest() *TaobaotradegetAPIRequest {
	return &TaobaotradegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradegetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaotradegetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotradegetAPIRequest) GetFields() string {
	return r._fields
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaotradegetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradegetAPIRequest) GetTid() int64 {
	return r._tid
}
