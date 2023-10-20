package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse 根据天猫id查询门店服务授权 API返回值
// alibaba.mallitemcenter.entitledservice.supplier.query
//
// 根据天猫id查询门店服务授权
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse struct {
	model.CommonResponse
	AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponseModel).Reset()
}

// AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponseModel is 根据天猫id查询门店服务授权 成功返回结果
type AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mallitemcenter_entitledservice_supplier_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 统一返回结果
	Result *AlibabaMallitemcenterEntitledserviceSupplierQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse)
	},
}

// GetAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse 从 sync.Pool 获取 AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse
func GetAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse() *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse {
	return poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse.Get().(*AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse)
}

// ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse 将 AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse(v *AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse) {
	v.Reset()
	poolAlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse.Put(v)
}
