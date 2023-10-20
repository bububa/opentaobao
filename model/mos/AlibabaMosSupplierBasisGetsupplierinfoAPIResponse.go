package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosSupplierBasisGetsupplierinfoAPIResponse 获取供应商基础信息 API返回值
// alibaba.mos.supplier.basis.getsupplierinfo
//
// 基于供应商id获取供应商基础脱敏信息
type AlibabaMosSupplierBasisGetsupplierinfoAPIResponse struct {
	model.CommonResponse
	AlibabaMosSupplierBasisGetsupplierinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosSupplierBasisGetsupplierinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosSupplierBasisGetsupplierinfoAPIResponseModel).Reset()
}

// AlibabaMosSupplierBasisGetsupplierinfoAPIResponseModel is 获取供应商基础信息 成功返回结果
type AlibabaMosSupplierBasisGetsupplierinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_supplier_basis_getsupplierinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *JsonResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosSupplierBasisGetsupplierinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosSupplierBasisGetsupplierinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosSupplierBasisGetsupplierinfoAPIResponse)
	},
}

// GetAlibabaMosSupplierBasisGetsupplierinfoAPIResponse 从 sync.Pool 获取 AlibabaMosSupplierBasisGetsupplierinfoAPIResponse
func GetAlibabaMosSupplierBasisGetsupplierinfoAPIResponse() *AlibabaMosSupplierBasisGetsupplierinfoAPIResponse {
	return poolAlibabaMosSupplierBasisGetsupplierinfoAPIResponse.Get().(*AlibabaMosSupplierBasisGetsupplierinfoAPIResponse)
}

// ReleaseAlibabaMosSupplierBasisGetsupplierinfoAPIResponse 将 AlibabaMosSupplierBasisGetsupplierinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosSupplierBasisGetsupplierinfoAPIResponse(v *AlibabaMosSupplierBasisGetsupplierinfoAPIResponse) {
	v.Reset()
	poolAlibabaMosSupplierBasisGetsupplierinfoAPIResponse.Put(v)
}
