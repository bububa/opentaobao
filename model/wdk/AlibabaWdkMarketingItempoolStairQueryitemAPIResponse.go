package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolstairqueryitemAPIResponse 换购商品查询 API返回值
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabawdkmarketingitempoolstairqueryitemAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingitempoolstairqueryitemAPIResponseModel
}

// AlibabawdkmarketingitempoolstairqueryitemAPIResponseModel is 换购商品查询 成功返回结果
type AlibabawdkmarketingitempoolstairqueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
