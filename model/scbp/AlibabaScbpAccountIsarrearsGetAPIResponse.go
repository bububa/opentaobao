package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountisarrearsgetAPIResponse 查询关键词推广账户是否欠款 API返回值
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
type AlibabascbpaccountisarrearsgetAPIResponse struct {
	model.CommonResponse
	AlibabascbpaccountisarrearsgetAPIResponseModel
}

// AlibabascbpaccountisarrearsgetAPIResponseModel is 查询关键词推广账户是否欠款 成功返回结果
type AlibabascbpaccountisarrearsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_isarrears_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 客户的关键词推广账户是否欠款
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
