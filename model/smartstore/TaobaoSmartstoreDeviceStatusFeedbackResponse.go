package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备在线状态回流 APIResponse
taobao.smartstore.device.status.feedback

智能硬件设备状态回流
*/
type TaobaoSmartstoreDeviceStatusFeedbackAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceStatusFeedbackResponse
}

type TaobaoSmartstoreDeviceStatusFeedbackResponse struct {
    XMLName xml.Name `xml:"smartstore_device_status_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
