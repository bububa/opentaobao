package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件互动云货架批量数据回流 API返回值 
taobao.smartstore.device.iashelf.batch.feedback

智慧门店互动云货架设备批量回流，
只能回流单个设备的批量业务数据
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
type TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceIashelfBatchFeedbackResponse
}

// 智能硬件互动云货架批量数据回流 成功返回结果
type TaobaoSmartstoreDeviceIashelfBatchFeedbackResponse struct {
    XMLName xml.Name `xml:"smartstore_device_iashelf_batch_feedback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回流结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
