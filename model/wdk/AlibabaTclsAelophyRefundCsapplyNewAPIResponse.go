package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundcsapplynewAPIResponse 代客退 API返回值
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
type AlibabatclsaelophyrefundcsapplynewAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophyrefundcsapplynewAPIResponseModel
}

// AlibabatclsaelophyrefundcsapplynewAPIResponseModel is 代客退 成功返回结果
type AlibabatclsaelophyrefundcsapplynewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_csapply_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabatclsaelophyrefundcsapplynewApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
