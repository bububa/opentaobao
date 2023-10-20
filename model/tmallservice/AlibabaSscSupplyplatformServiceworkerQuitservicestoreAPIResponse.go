package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse 工人退出网点 API返回值
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
type AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponseModel is 工人退出网点 成功返回结果
type AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_quitservicestore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceworkerQuitservicestoreResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse
func GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse() *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse {
	return poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse.Get().(*AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse 将 AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse(v *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse.Put(v)
}
