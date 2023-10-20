package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountdeleteactivityAPIResponse 删除商品特价活动 API返回值
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabahmmarketingitemdiscountdeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitemdiscountdeleteactivityAPIResponseModel
}

// AlibabahmmarketingitemdiscountdeleteactivityAPIResponseModel is 删除商品特价活动 成功返回结果
type AlibabahmmarketingitemdiscountdeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
