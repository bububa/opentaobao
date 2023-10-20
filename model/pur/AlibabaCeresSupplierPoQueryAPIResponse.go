package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCeresSupplierPoQueryAPIResponse 采购供应商订单查询接口 API返回值
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
type AlibabaCeresSupplierPoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaCeresSupplierPoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCeresSupplierPoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCeresSupplierPoQueryAPIResponseModel).Reset()
}

// AlibabaCeresSupplierPoQueryAPIResponseModel is 采购供应商订单查询接口 成功返回结果
type AlibabaCeresSupplierPoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ceres_supplier_po_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息体
	Result *AlibabaCeresSupplierPoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCeresSupplierPoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCeresSupplierPoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCeresSupplierPoQueryAPIResponse)
	},
}

// GetAlibabaCeresSupplierPoQueryAPIResponse 从 sync.Pool 获取 AlibabaCeresSupplierPoQueryAPIResponse
func GetAlibabaCeresSupplierPoQueryAPIResponse() *AlibabaCeresSupplierPoQueryAPIResponse {
	return poolAlibabaCeresSupplierPoQueryAPIResponse.Get().(*AlibabaCeresSupplierPoQueryAPIResponse)
}

// ReleaseAlibabaCeresSupplierPoQueryAPIResponse 将 AlibabaCeresSupplierPoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCeresSupplierPoQueryAPIResponse(v *AlibabaCeresSupplierPoQueryAPIResponse) {
	v.Reset()
	poolAlibabaCeresSupplierPoQueryAPIResponse.Put(v)
}
