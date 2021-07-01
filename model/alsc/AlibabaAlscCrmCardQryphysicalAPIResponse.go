package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardQryphysicalAPIResponse
查询物理卡 API返回值
alibaba.alsc.crm.card.qryphysical

查询物理卡 */
type AlibabaAlscCrmCardQryphysicalAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardQryphysicalAPIResponseModel
}

// AlibabaAlscCrmCardQryphysicalAPIResponseModel is 查询物理卡 成功返回结果
type AlibabaAlscCrmCardQryphysicalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_qryphysical_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
