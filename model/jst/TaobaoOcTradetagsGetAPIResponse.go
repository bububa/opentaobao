package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradetagsGetAPIResponse 根据订单查询订单标签 API返回值
// taobao.oc.tradetags.get
//
// 根据订单查询订单标签。&lt;br/&gt;
// 返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。&lt;br/&gt;
// 官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明&lt;br/&gt;
// 主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865&lt;br/&gt;
type TaobaoOcTradetagsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOcTradetagsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOcTradetagsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOcTradetagsGetAPIResponseModel).Reset()
}

// TaobaoOcTradetagsGetAPIResponseModel is 根据订单查询订单标签 成功返回结果
type TaobaoOcTradetagsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_tradetags_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TradeTags []TradeTagRelationDo `json:"trade_tags,omitempty" xml:"trade_tags>trade_tag_relation_do,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOcTradetagsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeTags = m.TradeTags[:0]
}

var poolTaobaoOcTradetagsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOcTradetagsGetAPIResponse)
	},
}

// GetTaobaoOcTradetagsGetAPIResponse 从 sync.Pool 获取 TaobaoOcTradetagsGetAPIResponse
func GetTaobaoOcTradetagsGetAPIResponse() *TaobaoOcTradetagsGetAPIResponse {
	return poolTaobaoOcTradetagsGetAPIResponse.Get().(*TaobaoOcTradetagsGetAPIResponse)
}

// ReleaseTaobaoOcTradetagsGetAPIResponse 将 TaobaoOcTradetagsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOcTradetagsGetAPIResponse(v *TaobaoOcTradetagsGetAPIResponse) {
	v.Reset()
	poolTaobaoOcTradetagsGetAPIResponse.Put(v)
}
