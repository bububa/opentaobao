package smartstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartstoredevicestatusfeedbackAPIResponse 设备在线状态回流 API返回值
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
type TaobaosmartstoredevicestatusfeedbackAPIResponse struct {
	model.CommonResponse
	TaobaosmartstoredevicestatusfeedbackAPIResponseModel
}

// TaobaosmartstoredevicestatusfeedbackAPIResponseModel is 设备在线状态回流 成功返回结果
type TaobaosmartstoredevicestatusfeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"smartstore_device_status_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
