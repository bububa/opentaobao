package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempooladditemAPIResponse 增加商品池里面的商品 API返回值
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabahmmarketingitempooladditemAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitempooladditemAPIResponseModel
}

// AlibabahmmarketingitempooladditemAPIResponseModel is 增加商品池里面的商品 成功返回结果
type AlibabahmmarketingitempooladditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
