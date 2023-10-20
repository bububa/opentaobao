package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse 商品池删除商品 API返回值
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
type AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel
}

// AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel is 商品池删除商品 成功返回结果
type AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
