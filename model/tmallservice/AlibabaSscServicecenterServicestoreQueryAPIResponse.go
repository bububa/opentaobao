package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscServicecenterServicestoreQueryAPIResponse 根据天猫id查询门店信息 API返回值
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
type AlibabaSscServicecenterServicestoreQueryAPIResponse struct {
	model.CommonResponse
	AlibabaSscServicecenterServicestoreQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscServicecenterServicestoreQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscServicecenterServicestoreQueryAPIResponseModel).Reset()
}

// AlibabaSscServicecenterServicestoreQueryAPIResponseModel is 根据天猫id查询门店信息 成功返回结果
type AlibabaSscServicecenterServicestoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_servicecenter_servicestore_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscServicecenterServicestoreQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscServicecenterServicestoreQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscServicecenterServicestoreQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscServicecenterServicestoreQueryAPIResponse)
	},
}

// GetAlibabaSscServicecenterServicestoreQueryAPIResponse 从 sync.Pool 获取 AlibabaSscServicecenterServicestoreQueryAPIResponse
func GetAlibabaSscServicecenterServicestoreQueryAPIResponse() *AlibabaSscServicecenterServicestoreQueryAPIResponse {
	return poolAlibabaSscServicecenterServicestoreQueryAPIResponse.Get().(*AlibabaSscServicecenterServicestoreQueryAPIResponse)
}

// ReleaseAlibabaSscServicecenterServicestoreQueryAPIResponse 将 AlibabaSscServicecenterServicestoreQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscServicecenterServicestoreQueryAPIResponse(v *AlibabaSscServicecenterServicestoreQueryAPIResponse) {
	v.Reset()
	poolAlibabaSscServicecenterServicestoreQueryAPIResponse.Put(v)
}
