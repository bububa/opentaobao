package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerSaveAPIResponse 服务商绑定工人 API返回值
// alibaba.ssc.supplyplatform.serviceworker.save
//
// 服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
type AlibabaSscSupplyplatformServiceworkerSaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceworkerSaveAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceworkerSaveAPIResponseModel is 服务商绑定工人 成功返回结果
type AlibabaSscSupplyplatformServiceworkerSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceworkerSaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceworkerSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceworkerSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceworkerSaveAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceworkerSaveAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerSaveAPIResponse
func GetAlibabaSscSupplyplatformServiceworkerSaveAPIResponse() *AlibabaSscSupplyplatformServiceworkerSaveAPIResponse {
	return poolAlibabaSscSupplyplatformServiceworkerSaveAPIResponse.Get().(*AlibabaSscSupplyplatformServiceworkerSaveAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerSaveAPIResponse 将 AlibabaSscSupplyplatformServiceworkerSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerSaveAPIResponse(v *AlibabaSscSupplyplatformServiceworkerSaveAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerSaveAPIResponse.Put(v)
}
