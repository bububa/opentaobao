package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradetagsGetAPIRequest 根据订单查询订单标签 API请求
// taobao.oc.tradetags.get
//
// 根据订单查询订单标签。<br/>
// 返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
// 官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
// 主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
type TaobaoOcTradetagsGetAPIRequest struct {
	model.Params
	// 交易主订单id
	_tid int64
	// 是否查询历史标签
	_history int64
	// 不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
	_tagTypes []string
	// 不填，则不做标签名称过滤
	_tagNames []string
}

// NewTaobaoOcTradetagsGetRequest 初始化TaobaoOcTradetagsGetAPIRequest对象
func NewTaobaoOcTradetagsGetRequest() *TaobaoOcTradetagsGetAPIRequest {
	return &TaobaoOcTradetagsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetagsGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.tradetags.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetagsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTid is Tid Setter
// 交易主订单id
func (r *TaobaoOcTradetagsGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOcTradetagsGetAPIRequest) GetTid() int64 {
	return r._tid
}

// SetHistory is History Setter
// 是否查询历史标签
func (r *TaobaoOcTradetagsGetAPIRequest) SetHistory(_history int64) error {
	r._history = _history
	r.Set("history", _history)
	return nil
}

// GetHistory History Getter
func (r TaobaoOcTradetagsGetAPIRequest) GetHistory() int64 {
	return r._history
}

// SetTagTypes is TagTypes Setter
// 不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
func (r *TaobaoOcTradetagsGetAPIRequest) SetTagTypes(_tagTypes []string) error {
	r._tagTypes = _tagTypes
	r.Set("tag_types", _tagTypes)
	return nil
}

// GetTagTypes TagTypes Getter
func (r TaobaoOcTradetagsGetAPIRequest) GetTagTypes() []string {
	return r._tagTypes
}

// SetTagNames is TagNames Setter
// 不填，则不做标签名称过滤
func (r *TaobaoOcTradetagsGetAPIRequest) SetTagNames(_tagNames []string) error {
	r._tagNames = _tagNames
	r.Set("tag_names", _tagNames)
	return nil
}

// GetTagNames TagNames Getter
func (r TaobaoOcTradetagsGetAPIRequest) GetTagNames() []string {
	return r._tagNames
}
