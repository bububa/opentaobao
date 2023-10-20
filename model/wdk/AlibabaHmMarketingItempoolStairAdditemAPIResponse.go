package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolstairadditemAPIResponse 商品池阶梯商品添加 API返回值
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabahmmarketingitempoolstairadditemAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitempoolstairadditemAPIResponseModel
}

// AlibabahmmarketingitempoolstairadditemAPIResponseModel is 商品池阶梯商品添加 成功返回结果
type AlibabahmmarketingitempoolstairadditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_stair_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
