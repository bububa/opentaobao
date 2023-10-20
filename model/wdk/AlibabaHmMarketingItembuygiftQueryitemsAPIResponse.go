package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitembuygiftqueryitemsAPIResponse 查询买赠活动下的商品 API返回值
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabahmmarketingitembuygiftqueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitembuygiftqueryitemsAPIResponseModel
}

// AlibabahmmarketingitembuygiftqueryitemsAPIResponseModel is 查询买赠活动下的商品 成功返回结果
type AlibabahmmarketingitembuygiftqueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
