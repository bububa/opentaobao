package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicegetAPIResponse 获取用户可用发票列表 API返回值
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
type AlitripbtripinvoicegetAPIResponse struct {
	model.CommonResponse
	AlitripbtripinvoicegetAPIResponseModel
}

// AlitripbtripinvoicegetAPIResponseModel is 获取用户可用发票列表 成功返回结果
type AlitripbtripinvoicegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
