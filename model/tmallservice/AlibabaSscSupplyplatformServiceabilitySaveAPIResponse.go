package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceabilitySaveAPIResponse 保存服务能力 API返回值
// alibaba.ssc.supplyplatform.serviceability.save
//
// 保存服务能力
type AlibabaSscSupplyplatformServiceabilitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceabilitySaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceabilitySaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceabilitySaveAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceabilitySaveAPIResponseModel is 保存服务能力 成功返回结果
type AlibabaSscSupplyplatformServiceabilitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceability_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceabilitySaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceabilitySaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceabilitySaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceabilitySaveAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceabilitySaveAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceabilitySaveAPIResponse
func GetAlibabaSscSupplyplatformServiceabilitySaveAPIResponse() *AlibabaSscSupplyplatformServiceabilitySaveAPIResponse {
	return poolAlibabaSscSupplyplatformServiceabilitySaveAPIResponse.Get().(*AlibabaSscSupplyplatformServiceabilitySaveAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceabilitySaveAPIResponse 将 AlibabaSscSupplyplatformServiceabilitySaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceabilitySaveAPIResponse(v *AlibabaSscSupplyplatformServiceabilitySaveAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceabilitySaveAPIResponse.Put(v)
}
