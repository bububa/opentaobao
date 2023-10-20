package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorInfoUploadAPIResponse 五道口仓库悬挂链信息上报 API返回值
// taobao.wdk.equipment.conveyor.info.upload
//
// 五道口仓库悬挂链信息上传
type TaobaoWdkEquipmentConveyorInfoUploadAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorInfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel is 五道口仓库悬挂链信息上报 成功返回结果
type TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_info_upload_response"`
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
func (m *TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.ServiceErrorCode = ""
	m.ServiceErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoWdkEquipmentConveyorInfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorInfoUploadAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorInfoUploadAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorInfoUploadAPIResponse
func GetTaobaoWdkEquipmentConveyorInfoUploadAPIResponse() *TaobaoWdkEquipmentConveyorInfoUploadAPIResponse {
	return poolTaobaoWdkEquipmentConveyorInfoUploadAPIResponse.Get().(*TaobaoWdkEquipmentConveyorInfoUploadAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorInfoUploadAPIResponse 将 TaobaoWdkEquipmentConveyorInfoUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorInfoUploadAPIResponse(v *TaobaoWdkEquipmentConveyorInfoUploadAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorInfoUploadAPIResponse.Put(v)
}
