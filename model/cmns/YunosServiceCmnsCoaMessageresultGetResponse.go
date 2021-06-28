package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息发送到达查询 APIResponse
yunos.service.cmns.coa.messageresult.get

CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
*/
type YunosServiceCmnsCoaMessageresultGetAPIResponse struct {
    model.CommonResponse
    // Response *YunosServiceCmnsCoaMessageresultGetResponse `json:"yunos_service_cmns_coa_messageresult_get_response,omitempty"` 
    YunosServiceCmnsCoaMessageresultGetResponse
}

/* model for simplify = false
type YunosServiceCmnsCoaMessageresultGetResponse struct {

    // 200表示查询成功
    
    Status   int64 `json:"status,omitempty"`
    

    // 接口查询出错提示信息
    
    Message   string `json:"message,omitempty"`
    

    // 具体的消息返回值
    
    Data  *struct {
        MessageResult  *MessageResult `json:"message_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type YunosServiceCmnsCoaMessageresultGetResponse struct {

    // 200表示查询成功
    Status   int64 `json:"status,omitempty"`

    // 接口查询出错提示信息
    Message   string `json:"message,omitempty"`

    // 具体的消息返回值
    Data   *MessageResult `json:"data,omitempty"`

}
