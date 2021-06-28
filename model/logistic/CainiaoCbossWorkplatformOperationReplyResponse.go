package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单操作回传 APIResponse
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformOperationReplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cboss_workplatform_operation_reply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CainiaoCbossWorkplatformOperationReplyResult `json:"result,omitempty" xml:"