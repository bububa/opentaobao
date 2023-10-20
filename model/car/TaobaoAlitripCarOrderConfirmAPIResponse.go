package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderConfirmAPIResponse 司机应答API API返回值
// taobao.alitrip.car.order.confirm
//
// 航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
type TaobaoAlitripCarOrderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderConfirmAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderConfirmAPIResponseModel is 司机应答API 成功返回结果
type TaobaoAlitripCarOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.MessageCode = 0
}

var poolTaobaoAlitripCarOrderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderConfirmAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderConfirmAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderConfirmAPIResponse
func GetTaobaoAlitripCarOrderConfirmAPIResponse() *TaobaoAlitripCarOrderConfirmAPIResponse {
	return poolTaobaoAlitripCarOrderConfirmAPIResponse.Get().(*TaobaoAlitripCarOrderConfirmAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderConfirmAPIResponse 将 TaobaoAlitripCarOrderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderConfirmAPIResponse(v *TaobaoAlitripCarOrderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderConfirmAPIResponse.Put(v)
}
