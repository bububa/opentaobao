package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRecommendCardinfoGetAPIResponse 快应用卡片信息 API返回值
// alibaba.alihealth.recommend.cardinfo.get
//
// 快应用卡片信息
type AlibabaAlihealthRecommendCardinfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthRecommendCardinfoGetAPIResponseModel
}

// AlibabaAlihealthRecommendCardinfoGetAPIResponseModel is 快应用卡片信息 成功返回结果
type AlibabaAlihealthRecommendCardinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_recommend_cardinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	CardResult *ServiceResult `json:"card_result,omitempty" xml:"card_result,omitempty"`
}
