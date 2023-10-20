package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarDriverStatusUpdateAPIResponse 司机服务状态更新接口 API返回值
// taobao.alitrip.car.driver.status.update
//
// 飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
type TaobaoAlitripCarDriverStatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarDriverStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarDriverStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarDriverStatusUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripCarDriverStatusUpdateAPIResponseModel is 司机服务状态更新接口 成功返回结果
type TaobaoAlitripCarDriverStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_driver_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 其它数据，预留，暂不使用
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarDriverStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Data = ""
	m.MessageCode = 0
}

var poolTaobaoAlitripCarDriverStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarDriverStatusUpdateAPIResponse)
	},
}

// GetTaobaoAlitripCarDriverStatusUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarDriverStatusUpdateAPIResponse
func GetTaobaoAlitripCarDriverStatusUpdateAPIResponse() *TaobaoAlitripCarDriverStatusUpdateAPIResponse {
	return poolTaobaoAlitripCarDriverStatusUpdateAPIResponse.Get().(*TaobaoAlitripCarDriverStatusUpdateAPIResponse)
}

// ReleaseTaobaoAlitripCarDriverStatusUpdateAPIResponse 将 TaobaoAlitripCarDriverStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarDriverStatusUpdateAPIResponse(v *TaobaoAlitripCarDriverStatusUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarDriverStatusUpdateAPIResponse.Put(v)
}
