package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopenposdiscountcodecreateAPIResponse pos一物一码创建 API返回值
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
type AlibabawdkmarketingopenposdiscountcodecreateAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingopenposdiscountcodecreateAPIResponseModel
}

// AlibabawdkmarketingopenposdiscountcodecreateAPIResponseModel is pos一物一码创建 成功返回结果
type AlibabawdkmarketingopenposdiscountcodecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_pos_discount_code_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
