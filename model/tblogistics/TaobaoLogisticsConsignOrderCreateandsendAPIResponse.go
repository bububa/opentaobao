package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsConsignOrderCreateandsendAPIResponse 创建订单并发货 API返回值
// taobao.logistics.consign.order.createandsend
//
// 创建物流订单，并发货。
type TaobaoLogisticsConsignOrderCreateandsendAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsConsignOrderCreateandsendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsConsignOrderCreateandsendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsConsignOrderCreateandsendAPIResponseModel).Reset()
}

// TaobaoLogisticsConsignOrderCreateandsendAPIResponseModel is 创建订单并发货 成功返回结果
type TaobaoLogisticsConsignOrderCreateandsendAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_consign_order_createandsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsConsignOrderCreateandsendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultDesc = ""
	m.OrderId = 0
	m.IsSuccess = false
}

var poolTaobaoLogisticsConsignOrderCreateandsendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsConsignOrderCreateandsendAPIResponse)
	},
}

// GetTaobaoLogisticsConsignOrderCreateandsendAPIResponse 从 sync.Pool 获取 TaobaoLogisticsConsignOrderCreateandsendAPIResponse
func GetTaobaoLogisticsConsignOrderCreateandsendAPIResponse() *TaobaoLogisticsConsignOrderCreateandsendAPIResponse {
	return poolTaobaoLogisticsConsignOrderCreateandsendAPIResponse.Get().(*TaobaoLogisticsConsignOrderCreateandsendAPIResponse)
}

// ReleaseTaobaoLogisticsConsignOrderCreateandsendAPIResponse 将 TaobaoLogisticsConsignOrderCreateandsendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsConsignOrderCreateandsendAPIResponse(v *TaobaoLogisticsConsignOrderCreateandsendAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsConsignOrderCreateandsendAPIResponse.Put(v)
}
