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
    YunosServiceCmnsCoaMessageAcksListResponse
}

type YunosServiceCmnsCoaMessageAcksListResponse struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_message_acks_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页结果对象
    
    Data   *PaginationQueryResult `json:"data,omitempty" xml:"data,omitempty"`

    
    // 接口查询出错提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 200表示查询成功
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
}
