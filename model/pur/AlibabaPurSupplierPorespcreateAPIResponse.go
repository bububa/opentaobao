package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierPorespcreateAPIResponse po反馈创建 API返回值
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
type AlibabaPurSupplierPorespcreateAPIResponse struct {
	model.CommonResponse
	AlibabaPurSupplierPorespcreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurSupplierPorespcreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurSupplierPorespcreateAPIResponseModel).Reset()
}

// AlibabaPurSupplierPorespcreateAPIResponseModel is po反馈创建 成功返回结果
type AlibabaPurSupplierPorespcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_supplier_porespcreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurSupplierPorespcreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurSupplierPorespcreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurSupplierPorespcreateAPIResponse)
	},
}

// GetAlibabaPurSupplierPorespcreateAPIResponse 从 sync.Pool 获取 AlibabaPurSupplierPorespcreateAPIResponse
func GetAlibabaPurSupplierPorespcreateAPIResponse() *AlibabaPurSupplierPorespcreateAPIResponse {
	return poolAlibabaPurSupplierPorespcreateAPIResponse.Get().(*AlibabaPurSupplierPorespcreateAPIResponse)
}

// ReleaseAlibabaPurSupplierPorespcreateAPIResponse 将 AlibabaPurSupplierPorespcreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurSupplierPorespcreateAPIResponse(v *AlibabaPurSupplierPorespcreateAPIResponse) {
	v.Reset()
	poolAlibabaPurSupplierPorespcreateAPIResponse.Put(v)
}
