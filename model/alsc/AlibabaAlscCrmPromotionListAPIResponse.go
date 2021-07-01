package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPromotionListAPIResponse
获取促销规则列表 API返回值
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表 */
type AlibabaAlscCrmPromotionListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPromotionListAPIResponseModel
}

// AlibabaAlscCrmPromotionListAPIResponseModel is 获取促销规则列表 成功返回结果
type AlibabaAlscCrmPromotionListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_promotion_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
