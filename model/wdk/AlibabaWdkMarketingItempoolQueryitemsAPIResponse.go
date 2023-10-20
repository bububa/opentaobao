package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolqueryitemsAPIResponse 查询商品池活动下的商品 API返回值
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabawdkmarketingitempoolqueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingitempoolqueryitemsAPIResponseModel
}

// AlibabawdkmarketingitempoolqueryitemsAPIResponseModel is 查询商品池活动下的商品 成功返回结果
type AlibabawdkmarketingitempoolqueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
