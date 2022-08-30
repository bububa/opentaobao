package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCustomerFindCustomerInfoAPIResponse 查询客户信息 API返回值
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
type AlibabaScbpAdCustomerFindCustomerInfoAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCustomerFindCustomerInfoAPIResponseModel
}

// AlibabaScbpAdCustomerFindCustomerInfoAPIResponseModel is 查询客户信息 成功返回结果
type AlibabaScbpAdCustomerFindCustomerInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_customer_find_customer_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回类
	Result *AlibabaScbpAdCustomerFindCustomerInfoResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
