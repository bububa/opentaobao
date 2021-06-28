package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息撤回 APIResponse
yunos.service.cmns.coa.message.cancel

此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
*/
type YunosServiceCmnsCoaMessageCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_message_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回状态 200表示调用成功
    
    Status   int64 `json:"status,omitempty" xml:"