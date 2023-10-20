package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyvoucherqueryAPIResponse 查询单个代金券信息 API返回值
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
type AlitripmerchantgalaxyvoucherqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyvoucherqueryAPIResponseModel
}

// AlitripmerchantgalaxyvoucherqueryAPIResponseModel is 查询单个代金券信息 成功返回结果
type AlitripmerchantgalaxyvoucherqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_voucher_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *AlitripmerchantgalaxyvoucherqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
