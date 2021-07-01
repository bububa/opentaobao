package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息详情查询 API返回值 
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
*/
type YunosServiceCmnsCoaMessageGetAPIResponse struct {
    model.CommonResponse
    YunosServiceCmnsCoaMessageGetAPIResponseModel
}

// 消息详情查询 成功返回结果
type YunosServiceCmnsCoaMessageGetAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_message_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 消息内容
    Data   *MessageDetailResult `json:"data,omitempty" xml:"data,omitempty"`
    // 接口查询出错提示信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 200表示查询成功
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
