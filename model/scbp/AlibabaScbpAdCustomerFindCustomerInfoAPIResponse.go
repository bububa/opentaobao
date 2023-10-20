package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcustomerfindcustomerinfoAPIResponse 查询客户信息 API返回值
// alibaba.scbp.ad.customer.find.customer.info
//
// 查询客户信息
type AlibabascbpadcustomerfindcustomerinfoAPIResponse struct {
	model.CommonResponse
	AlibabascbpadcustomerfindcustomerinfoAPIResponseModel
}

// AlibabascbpadcustomerfindcustomerinfoAPIResponseModel is 查询客户信息 成功返回结果
type AlibabascbpadcustomerfindcustomerinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_customer_find_customer_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回类
	Result *AlibabascbpadcustomerfindcustomerinfoResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
