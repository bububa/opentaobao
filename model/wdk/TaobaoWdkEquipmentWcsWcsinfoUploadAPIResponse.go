package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse 悬挂链业务信息上传 API返回值
// taobao.wdk.equipment.wcs.wcsinfo.upload
//
// 五道口仓库悬挂链信息上传
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponseModel is 悬挂链业务信息上传 成功返回结果
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_wcs_wcsinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errorCode
	ServiceErrorCode string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
	// errorMsg
	ServiceErrorMsg string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.ServiceErrorCode = ""
	m.ServiceErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse)
	},
}

// GetTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse
func GetTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse() *TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse {
	return poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse.Get().(*TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse)
}

// ReleaseTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse 将 TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse(v *TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse.Put(v)
}
