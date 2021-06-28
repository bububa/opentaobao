package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息发送到达查询 APIResponse
yunos.service.cmns.coa.messageresult.get

CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
*/
type YunosServiceCmnsCoaMessageresultGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_messageresult_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 200表示查询成功
    
    Status   int64 `json:"status,omitempty" xml:"