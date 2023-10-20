package larkiot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderGetgoodslistAPIResponse iot渠道获取卖品信息 API返回值
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
type TaobaoLarkIotOrderGetgoodslistAPIResponse struct {
	model.CommonResponse
	TaobaoLarkIotOrderGetgoodslistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLarkIotOrderGetgoodslistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLarkIotOrderGetgoodslistAPIResponseModel).Reset()
}

// TaobaoLarkIotOrderGetgoodslistAPIResponseModel is iot渠道获取卖品信息 成功返回结果
type TaobaoLarkIotOrderGetgoodslistAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_iot_order_getgoodslist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖品信息列表
	Data *BizSingleResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLarkIotOrderGetgoodslistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoLarkIotOrderGetgoodslistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLarkIotOrderGetgoodslistAPIResponse)
	},
}

// GetTaobaoLarkIotOrderGetgoodslistAPIResponse 从 sync.Pool 获取 TaobaoLarkIotOrderGetgoodslistAPIResponse
func GetTaobaoLarkIotOrderGetgoodslistAPIResponse() *TaobaoLarkIotOrderGetgoodslistAPIResponse {
	return poolTaobaoLarkIotOrderGetgoodslistAPIResponse.Get().(*TaobaoLarkIotOrderGetgoodslistAPIResponse)
}

// ReleaseTaobaoLarkIotOrderGetgoodslistAPIResponse 将 TaobaoLarkIotOrderGetgoodslistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLarkIotOrderGetgoodslistAPIResponse(v *TaobaoLarkIotOrderGetgoodslistAPIResponse) {
	v.Reset()
	poolTaobaoLarkIotOrderGetgoodslistAPIResponse.Put(v)
}
