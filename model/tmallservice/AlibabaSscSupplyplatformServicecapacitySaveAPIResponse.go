package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicecapacitySaveAPIResponse 保存服务容量 API返回值
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
type AlibabaSscSupplyplatformServicecapacitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServicecapacitySaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicecapacitySaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServicecapacitySaveAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServicecapacitySaveAPIResponseModel is 保存服务容量 成功返回结果
type AlibabaSscSupplyplatformServicecapacitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicecapacity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServicecapacitySaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServicecapacitySaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServicecapacitySaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServicecapacitySaveAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServicecapacitySaveAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServicecapacitySaveAPIResponse
func GetAlibabaSscSupplyplatformServicecapacitySaveAPIResponse() *AlibabaSscSupplyplatformServicecapacitySaveAPIResponse {
	return poolAlibabaSscSupplyplatformServicecapacitySaveAPIResponse.Get().(*AlibabaSscSupplyplatformServicecapacitySaveAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServicecapacitySaveAPIResponse 将 AlibabaSscSupplyplatformServicecapacitySaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicecapacitySaveAPIResponse(v *AlibabaSscSupplyplatformServicecapacitySaveAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicecapacitySaveAPIResponse.Put(v)
}
