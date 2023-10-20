package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse 工人取消请假 API返回值
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel is 工人取消请假 成功返回结果
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_cancelleave_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceworkerCancelleaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse
func GetAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse() *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse {
	return poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse.Get().(*AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse 将 AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse(v *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse.Put(v)
}
