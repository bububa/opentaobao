package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse 服务容量删除 API返回值
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
type AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServicecapacityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServicecapacityDeleteAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServicecapacityDeleteAPIResponseModel is 服务容量删除 成功返回结果
type AlibabaSscSupplyplatformServicecapacityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicecapacity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServicecapacityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicecapacityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse
func GetAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse() *AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse {
	return poolAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse.Get().(*AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse 将 AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse(v *AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicecapacityDeleteAPIResponse.Put(v)
}
