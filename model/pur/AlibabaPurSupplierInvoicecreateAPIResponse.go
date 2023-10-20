package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapursupplierinvoicecreateAPIResponse preInvoice创建 API返回值
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
type AlibabapursupplierinvoicecreateAPIResponse struct {
	model.CommonResponse
	AlibabapursupplierinvoicecreateAPIResponseModel
}

// AlibabapursupplierinvoicecreateAPIResponseModel is preInvoice创建 成功返回结果
type AlibabapursupplierinvoicecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_supplier_invoicecreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
