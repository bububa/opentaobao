package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeGetAPIRequest 获取单笔交易的部分信息(性能高) API请求
// taobao.trade.get
//
// 获取单笔交易的部分信息
// &lt;br/&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeGetAPIRequest) Reset() {
	r._fields = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTradeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeGetRequest()
	},
}

// GetTaobaoTradeGetRequest 从 sync.Pool 获取 TaobaoTradeGetAPIRequest
func GetTaobaoTradeGetAPIRequest() *TaobaoTradeGetAPIRequest {
	return poolTaobaoTradeGetAPIRequest.Get().(*TaobaoTradeGetAPIRequest)
}

// ReleaseTaobaoTradeGetAPIRequest 将 TaobaoTradeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeGetAPIRequest(v *TaobaoTradeGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeGetAPIRequest.Put(v)
}
