package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingdiscountitemaddasyncAPIResponse 特价批量新增商品 API返回值
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabahmmarketingdiscountitemaddasyncAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingdiscountitemaddasyncAPIResponseModel
}

// AlibabahmmarketingdiscountitemaddasyncAPIResponseModel is 特价批量新增商品 成功返回结果
type AlibabahmmarketingdiscountitemaddasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_discount_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
