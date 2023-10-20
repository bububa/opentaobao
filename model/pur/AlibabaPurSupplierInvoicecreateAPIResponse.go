package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierInvoicecreateAPIResponse preInvoice创建 API返回值
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
type AlibabaPurSupplierInvoicecreateAPIResponse struct {
	model.CommonResponse
	AlibabaPurSupplierInvoicecreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurSupplierInvoicecreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurSupplierInvoicecreateAPIResponseModel).Reset()
}

// AlibabaPurSupplierInvoicecreateAPIResponseModel is preInvoice创建 成功返回结果
type AlibabaPurSupplierInvoicecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_supplier_invoicecreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurSupplierInvoicecreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurSupplierInvoicecreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurSupplierInvoicecreateAPIResponse)
	},
}

// GetAlibabaPurSupplierInvoicecreateAPIResponse 从 sync.Pool 获取 AlibabaPurSupplierInvoicecreateAPIResponse
func GetAlibabaPurSupplierInvoicecreateAPIResponse() *AlibabaPurSupplierInvoicecreateAPIResponse {
	return poolAlibabaPurSupplierInvoicecreateAPIResponse.Get().(*AlibabaPurSupplierInvoicecreateAPIResponse)
}

// ReleaseAlibabaPurSupplierInvoicecreateAPIResponse 将 AlibabaPurSupplierInvoicecreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurSupplierInvoicecreateAPIResponse(v *AlibabaPurSupplierInvoicecreateAPIResponse) {
	v.Reset()
	poolAlibabaPurSupplierInvoicecreateAPIResponse.Put(v)
}
