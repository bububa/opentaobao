package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息详情查询 APIResponse
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
*/
type YunosServiceCmnsCoaMessageGetAPIResponse struct {
    model.CommonResponse
    Response *YunosServiceCmnsCoaMessageGetResponse `json:"yunos_service_cmns_coa_message_get_response,omitempty"`
}

type YunosServiceCmnsCoaMessageGetResponse struct {

    // 消息内容
    Data   *MessageDetailResult `json:"data,omitempty"`

    // 接口查询出错提示信息
    Message   string `json:"message,omitempty"`

    // 200表示查询成功
    Status   int64 `json:"status,omitempty"`

}
