package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单操作回传 APIResponse
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformOperationReplyAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCbossWorkplatformOperationReplyResponse `json:"cainiao_cboss_workplatform_operation_reply_response,omitempty"`
}

type CainiaoCbossWorkplatformOperationReplyResponse struct {

    // result
    Result   *CainiaoCbossWorkplatformOperationReplyResult `json:"result,omitempty"`

}
