package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送母婴提醒 API返回值 
alibaba.alihealth.baby.remind.batch.send

批量发送母婴提醒
*/
type AlibabaAlihealthBabyRemindBatchSendAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBabyRemindBatchSendAPIResponseModel
}

// 批量发送母婴提醒 成功返回结果
type AlibabaAlihealthBabyRemindBatchSendAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_baby_remind_batch_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
