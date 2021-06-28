package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息ack记录查询 APIResponse
yunos.service.cmns.coa.message.acks.list

第三方应用开发者调用此接口查询消息ack记录
*/
type YunosServiceCmnsCoaMessageAcksListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_message_acks_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页结果对象
    
    Data   *PaginationQueryResult `json:"data,omitempty" xml:"