package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeGetAPIRequest 获取单笔交易的部分信息(性能高) API请求
// taobao.trade.get
//
// 获取单笔交易的部分信息
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
type TaobaoTradeGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// NewTaobaoTradeGetRequest 初始化TaobaoTradeGetAPIRequest对象
func NewTaobaoTradeGetRequest() *TaobaoTradeGetAPIRequest {
	return &TaobaoTradeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
func (r *TaobaoTradeGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTradeGetAPIRequest) GetFields() string {
	return r._fields
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaoTradeGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeGetAPIRequest) GetTid() int64 {
	return r._tid
}
