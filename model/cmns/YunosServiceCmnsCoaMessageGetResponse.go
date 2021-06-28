package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息详情查询 APIResponse
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
*/
type YunosServiceCmnsCoaMessageGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_message_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息内容
    
    Data   *MessageDetailResult `json:"data,omitempty" xml:"