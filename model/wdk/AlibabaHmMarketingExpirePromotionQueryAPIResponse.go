package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingexpirepromotionqueryAPIResponse 短保优惠查询 API返回值
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabahmmarketingexpirepromotionqueryAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingexpirepromotionqueryAPIResponseModel
}

// AlibabahmmarketingexpirepromotionqueryAPIResponseModel is 短保优惠查询 成功返回结果
type AlibabahmmarketingexpirepromotionqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
