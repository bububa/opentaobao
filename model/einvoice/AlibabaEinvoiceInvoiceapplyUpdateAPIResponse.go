package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceinvoiceapplyupdateAPIResponse 商家开票申请单状态回传 API返回值
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
type AlibabaeinvoiceinvoiceapplyupdateAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoiceinvoiceapplyupdateAPIResponseModel
}

// AlibabaeinvoiceinvoiceapplyupdateAPIResponseModel is 商家开票申请单状态回传 成功返回结果
type AlibabaeinvoiceinvoiceapplyupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_invoiceapply_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 更新结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
