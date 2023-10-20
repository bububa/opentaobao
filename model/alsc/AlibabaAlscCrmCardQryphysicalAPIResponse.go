package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardqryphysicalAPIResponse 查询物理卡 API返回值
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
type AlibabaalsccrmcardqryphysicalAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmcardqryphysicalAPIResponseModel
}

// AlibabaalsccrmcardqryphysicalAPIResponseModel is 查询物理卡 成功返回结果
type AlibabaalsccrmcardqryphysicalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_qryphysical_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
