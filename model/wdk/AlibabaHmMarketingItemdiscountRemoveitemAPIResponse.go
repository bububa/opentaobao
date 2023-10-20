package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountremoveitemAPIResponse 移除报名的商品 API返回值
// alibaba.hm.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
type AlibabahmmarketingitemdiscountremoveitemAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitemdiscountremoveitemAPIResponseModel
}

// AlibabahmmarketingitemdiscountremoveitemAPIResponseModel is 移除报名的商品 成功返回结果
type AlibabahmmarketingitemdiscountremoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
