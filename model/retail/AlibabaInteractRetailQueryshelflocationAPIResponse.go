package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractRetailQueryshelflocationAPIResponse 查询货架和位置数据 API返回值
// alibaba.interact.retail.queryshelflocation
//
// 查询货架和位置数据
type AlibabaInteractRetailQueryshelflocationAPIResponse struct {
	model.CommonResponse
	AlibabaInteractRetailQueryshelflocationAPIResponseModel
}

// AlibabaInteractRetailQueryshelflocationAPIResponseModel is 查询货架和位置数据 成功返回结果
type AlibabaInteractRetailQueryshelflocationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_retail_queryshelflocation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaInteractRetailQueryshelflocationResult `json:"result,omitempty" xml:"result,omitempty"`
}
