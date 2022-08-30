package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketConsumeAPIResponse 天猫汽车二轮车电子凭证核销 API返回值
// tmall.aliauto.eticket.consume
//
// 天猫汽车二轮车行业门店电子凭证核销
type TmallAliautoEticketConsumeAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketConsumeAPIResponseModel
}

// TmallAliautoEticketConsumeAPIResponseModel is 天猫汽车二轮车电子凭证核销 成功返回结果
type TmallAliautoEticketConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
