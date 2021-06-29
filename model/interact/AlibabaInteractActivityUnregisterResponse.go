package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动关闭服务 APIResponse
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务
*/
type AlibabaInteractActivityUnregisterAPIResponse struct {
    model.CommonResponse
    AlibabaInteractActivityUnregisterResponse
}

type AlibabaInteractActivityUnregisterResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_activity_unregister_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 关闭活动结果
    
    UnregisterResult   *AllsparkResult `json:"unregister_result,omitempty" xml:"unregister_result,omitempty"`

    
}
