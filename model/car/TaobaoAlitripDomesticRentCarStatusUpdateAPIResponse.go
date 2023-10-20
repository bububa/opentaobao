package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse 航旅国内租车订单状态更新 API返回值
// taobao.alitrip.domestic.rent.car.status.update
//
// 航旅国内租车订单状态更新
type TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripDomesticRentCarStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripDomesticRentCarStatusUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripDomesticRentCarStatusUpdateAPIResponseModel is 航旅国内租车订单状态更新 成功返回结果
type TaobaoAlitripDomesticRentCarStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_domestic_rent_car_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 其它数据，预留，暂不使用
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码.code为0时表示成功
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripDomesticRentCarStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.MessageCode = 0
}

var poolTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse)
	},
}

// GetTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse
func GetTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse() *TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse {
	return poolTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse.Get().(*TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse)
}

// ReleaseTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse 将 TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse(v *TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripDomesticRentCarStatusUpdateAPIResponse.Put(v)
}
