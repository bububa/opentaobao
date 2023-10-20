package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscbusinessservicepricequeryAPIResponse 苹果查询服务供给报价 API返回值
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
type AlibabasscbusinessservicepricequeryAPIResponse struct {
	model.CommonResponse
	AlibabasscbusinessservicepricequeryAPIResponseModel
}

// AlibabasscbusinessservicepricequeryAPIResponseModel is 苹果查询服务供给报价 成功返回结果
type AlibabasscbusinessservicepricequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_business_serviceprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabasscbusinessservicepricequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
