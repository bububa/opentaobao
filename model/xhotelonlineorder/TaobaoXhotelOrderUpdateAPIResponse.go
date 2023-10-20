package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderUpdateAPIResponse 酒店订单发货接口 API返回值
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
type TaobaoXhotelOrderUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelOrderUpdateAPIResponseModel is 酒店订单发货接口 成功返回结果
type TaobaoXhotelOrderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回提示信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderUpdateAPIResponse)
	},
}

// GetTaobaoXhotelOrderUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderUpdateAPIResponse
func GetTaobaoXhotelOrderUpdateAPIResponse() *TaobaoXhotelOrderUpdateAPIResponse {
	return poolTaobaoXhotelOrderUpdateAPIResponse.Get().(*TaobaoXhotelOrderUpdateAPIResponse)
}

// ReleaseTaobaoXhotelOrderUpdateAPIResponse 将 TaobaoXhotelOrderUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderUpdateAPIResponse(v *TaobaoXhotelOrderUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderUpdateAPIResponse.Put(v)
}
