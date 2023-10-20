package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse 服务商添加工人 API返回值
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
type AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel is 服务商添加工人 成功返回结果
type AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用响应model
	Result *AlibabaSscSupplyplatformServiceworkerRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse
func GetAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse() *AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse {
	return poolAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse.Get().(*AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse 将 AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse(v *AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerRegisterAPIResponse.Put(v)
}
