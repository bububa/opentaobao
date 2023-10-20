package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicestoreOfflineAPIResponse 网点下线 API返回值
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
type AlibabaSscSupplyplatformServicestoreOfflineAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServicestoreOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicestoreOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServicestoreOfflineAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServicestoreOfflineAPIResponseModel is 网点下线 成功返回结果
type AlibabaSscSupplyplatformServicestoreOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicestore_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServicestoreOfflineResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicestoreOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServicestoreOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicestoreOfflineAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServicestoreOfflineAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServicestoreOfflineAPIResponse
func GetAlibabaSscSupplyplatformServicestoreOfflineAPIResponse() *AlibabaSscSupplyplatformServicestoreOfflineAPIResponse {
	return poolAlibabaSscSupplyplatformServicestoreOfflineAPIResponse.Get().(*AlibabaSscSupplyplatformServicestoreOfflineAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServicestoreOfflineAPIResponse 将 AlibabaSscSupplyplatformServicestoreOfflineAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicestoreOfflineAPIResponse(v *AlibabaSscSupplyplatformServicestoreOfflineAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicestoreOfflineAPIResponse.Put(v)
}
