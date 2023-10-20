package smartstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartstoreDeviceStatusFeedbackAPIResponse 设备在线状态回流 API返回值
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
type TaobaoSmartstoreDeviceStatusFeedbackAPIResponse struct {
	model.CommonResponse
	TaobaoSmartstoreDeviceStatusFeedbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartstoreDeviceStatusFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartstoreDeviceStatusFeedbackAPIResponseModel).Reset()
}

// TaobaoSmartstoreDeviceStatusFeedbackAPIResponseModel is 设备在线状态回流 成功返回结果
type TaobaoSmartstoreDeviceStatusFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"smartstore_device_status_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartstoreDeviceStatusFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoSmartstoreDeviceStatusFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartstoreDeviceStatusFeedbackAPIResponse)
	},
}

// GetTaobaoSmartstoreDeviceStatusFeedbackAPIResponse 从 sync.Pool 获取 TaobaoSmartstoreDeviceStatusFeedbackAPIResponse
func GetTaobaoSmartstoreDeviceStatusFeedbackAPIResponse() *TaobaoSmartstoreDeviceStatusFeedbackAPIResponse {
	return poolTaobaoSmartstoreDeviceStatusFeedbackAPIResponse.Get().(*TaobaoSmartstoreDeviceStatusFeedbackAPIResponse)
}

// ReleaseTaobaoSmartstoreDeviceStatusFeedbackAPIResponse 将 TaobaoSmartstoreDeviceStatusFeedbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartstoreDeviceStatusFeedbackAPIResponse(v *TaobaoSmartstoreDeviceStatusFeedbackAPIResponse) {
	v.Reset()
	poolTaobaoSmartstoreDeviceStatusFeedbackAPIResponse.Put(v)
}
