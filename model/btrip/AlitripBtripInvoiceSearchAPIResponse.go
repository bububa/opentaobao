package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSearchAPIResponse 根据发票抬头搜索发票 API返回值
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
type AlitripBtripInvoiceSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSearchAPIResponseModel
}

// AlitripBtripInvoiceSearchAPIResponseModel is 根据发票抬头搜索发票 成功返回结果
type AlitripBtripInvoiceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
