package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolitemaddasyncAPIResponse 商品池新增商品 API返回值
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
type AlibabahmmarketingitempoolitemaddasyncAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitempoolitemaddasyncAPIResponseModel
}

// AlibabahmmarketingitempoolitemaddasyncAPIResponseModel is 商品池新增商品 成功返回结果
type AlibabahmmarketingitempoolitemaddasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
