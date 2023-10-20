package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingpriceAPIResponse 促销价签服务 API返回值
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
type AlibabawdkmarketingpriceAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingpriceAPIResponseModel
}

// AlibabawdkmarketingpriceAPIResponseModel is 促销价签服务 成功返回结果
type AlibabawdkmarketingpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PromotionPriceResult `json:"result,omitempty" xml:"result,omitempty"`
}
