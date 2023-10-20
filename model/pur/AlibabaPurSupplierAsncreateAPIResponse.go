package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurSupplierAsncreateAPIResponse asn创建 API返回值
// alibaba.pur.supplier.asncreate
//
// asn创建
type AlibabaPurSupplierAsncreateAPIResponse struct {
	model.CommonResponse
	AlibabaPurSupplierAsncreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurSupplierAsncreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurSupplierAsncreateAPIResponseModel).Reset()
}

// AlibabaPurSupplierAsncreateAPIResponseModel is asn创建 成功返回结果
type AlibabaPurSupplierAsncreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_supplier_asncreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurSupplierAsncreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurSupplierAsncreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurSupplierAsncreateAPIResponse)
	},
}

// GetAlibabaPurSupplierAsncreateAPIResponse 从 sync.Pool 获取 AlibabaPurSupplierAsncreateAPIResponse
func GetAlibabaPurSupplierAsncreateAPIResponse() *AlibabaPurSupplierAsncreateAPIResponse {
	return poolAlibabaPurSupplierAsncreateAPIResponse.Get().(*AlibabaPurSupplierAsncreateAPIResponse)
}

// ReleaseAlibabaPurSupplierAsncreateAPIResponse 将 AlibabaPurSupplierAsncreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurSupplierAsncreateAPIResponse(v *AlibabaPurSupplierAsncreateAPIResponse) {
	v.Reset()
	poolAlibabaPurSupplierAsncreateAPIResponse.Put(v)
}
