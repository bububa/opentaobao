package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送母婴提醒 APIResponse
alibaba.alihealth.baby.remind.batch.send

批量发送母婴提醒
*/
type AlibabaAlihealthBabyRemindBatchSendAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBabyRemindBatchSendResponse
}

type AlibabaAlihealthBabyRemindBatchSendResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_baby_remind_batch_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
