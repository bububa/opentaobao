package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolDeleteactivityAPIResponse 删除商品池活动 API返回值
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaHmMarketingItempoolDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel
}

// AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel is 删除商品池活动 成功返回结果
type AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
