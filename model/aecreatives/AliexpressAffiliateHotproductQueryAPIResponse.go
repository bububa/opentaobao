package aecreatives

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateHotproductQueryAPIResponse
查询联盟爆品数据 API返回值
aliexpress.affiliate.hotproduct.query

查询联盟爆品API */
type AliexpressAffiliateHotproductQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateHotproductQueryAPIResponseModel
}

// AliexpressAffiliateHotproductQueryAPIResponseModel is 查询联盟爆品数据 成功返回结果
type AliexpressAffiliateHotproductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
