package cmns

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoamessageresultgetAPIResponse CMNS消息发送到达查询 API返回值
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
type YunosservicecmnscoamessageresultgetAPIResponse struct {
	model.CommonResponse
	YunosservicecmnscoamessageresultgetAPIResponseModel
}

// YunosservicecmnscoamessageresultgetAPIResponseModel is CMNS消息发送到达查询 成功返回结果
type YunosservicecmnscoamessageresultgetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_messageresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口查询出错提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 200表示查询成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 具体的消息返回值
	Data *MessageResult `json:"data,omitempty" xml:"data,omitempty"`
}
