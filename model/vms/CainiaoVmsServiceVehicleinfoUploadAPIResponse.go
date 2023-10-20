package vms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoVmsServiceVehicleinfoUploadAPIResponse 新能源车--外部车辆信息回传 API返回值
// cainiao.vms.service.vehicleinfo.upload
//
// 新能源车--外部车辆信息回传
type CainiaoVmsServiceVehicleinfoUploadAPIResponse struct {
	model.CommonResponse
	CainiaoVmsServiceVehicleinfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoVmsServiceVehicleinfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoVmsServiceVehicleinfoUploadAPIResponseModel).Reset()
}

// CainiaoVmsServiceVehicleinfoUploadAPIResponseModel is 新能源车--外部车辆信息回传 成功返回结果
type CainiaoVmsServiceVehicleinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_vms_service_vehicleinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AppBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoVmsServiceVehicleinfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoVmsServiceVehicleinfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoVmsServiceVehicleinfoUploadAPIResponse)
	},
}

// GetCainiaoVmsServiceVehicleinfoUploadAPIResponse 从 sync.Pool 获取 CainiaoVmsServiceVehicleinfoUploadAPIResponse
func GetCainiaoVmsServiceVehicleinfoUploadAPIResponse() *CainiaoVmsServiceVehicleinfoUploadAPIResponse {
	return poolCainiaoVmsServiceVehicleinfoUploadAPIResponse.Get().(*CainiaoVmsServiceVehicleinfoUploadAPIResponse)
}

// ReleaseCainiaoVmsServiceVehicleinfoUploadAPIResponse 将 CainiaoVmsServiceVehicleinfoUploadAPIResponse 保存到 sync.Pool
func ReleaseCainiaoVmsServiceVehicleinfoUploadAPIResponse(v *CainiaoVmsServiceVehicleinfoUploadAPIResponse) {
	v.Reset()
	poolCainiaoVmsServiceVehicleinfoUploadAPIResponse.Put(v)
}
