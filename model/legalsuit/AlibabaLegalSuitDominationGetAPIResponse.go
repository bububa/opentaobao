package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitDominationGetAPIResponse 查询管辖信息 API返回值
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
type AlibabaLegalSuitDominationGetAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitDominationGetAPIResponseModel
}

// AlibabaLegalSuitDominationGetAPIResponseModel is 查询管辖信息 成功返回结果
type AlibabaLegalSuitDominationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_domination_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
