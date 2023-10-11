package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolRemoveitemAPIResponse 移除商品池里面的商品 API返回值
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaHmMarketingItempoolRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolRemoveitemAPIResponseModel
}

// AlibabaHmMarketingItempoolRemoveitemAPIResponseModel is 移除商品池里面的商品 成功返回结果
type AlibabaHmMarketingItempoolRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
