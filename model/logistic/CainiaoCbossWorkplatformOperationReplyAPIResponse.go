package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单操作回传 API返回值 
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformOperationReplyAPIResponse struct {
    model.CommonResponse
    CainiaoCbossWorkplatformOperationReplyAPIResponseModel
}

// 菜鸟工单操作回传 成功返回结果
type CainiaoCbossWorkplatformOperationReplyAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_cboss_workplatform_operation_reply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CainiaoCbossWorkplatformOperationReplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
