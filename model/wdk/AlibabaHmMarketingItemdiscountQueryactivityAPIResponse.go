package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountqueryactivityAPIResponse 查找特价活动 API返回值
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabahmmarketingitemdiscountqueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitemdiscountqueryactivityAPIResponseModel
}

// AlibabahmmarketingitemdiscountqueryactivityAPIResponseModel is 查找特价活动 成功返回结果
type AlibabahmmarketingitemdiscountqueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询特价活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
