package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolexcludeskucodeAPIResponse 商品池排除商品【品类优惠使用】 API返回值
// alibaba.wdk.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
type AlibabawdkmarketingitempoolexcludeskucodeAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingitempoolexcludeskucodeAPIResponseModel
}

// AlibabawdkmarketingitempoolexcludeskucodeAPIResponseModel is 商品池排除商品【品类优惠使用】 成功返回结果
type AlibabawdkmarketingitempoolexcludeskucodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_excludeskucode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
