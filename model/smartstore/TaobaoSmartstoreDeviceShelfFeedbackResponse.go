package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件云货架数据回流 APIResponse
taobao.smartstore.device.shelf.feedback

智慧门店云货架设备回流
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
type TaobaoSmartstoreDeviceShelfFeedbackAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceShelfFeedbackResponse
}

type TaobaoSmartstoreDeviceShelfFeedbackResponse struct {
    XMLName xml.Name `xml:"smartstore_device_shelf_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回流结果
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
