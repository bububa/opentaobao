package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse 容器被悬挂链扫描 API返回值
// taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor
//
// 容器被悬挂链扫描
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel is 容器被悬挂链扫描 成功返回结果
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_wcsbtoc_containerscannedbyconveyor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// list
	List []WcsContainerScannedByConveyorDto `json:"list,omitempty" xml:"list>wcs_container_scanned_by_conveyor_dto,omitempty"`
	// errorCode
	ServiceErrorCode string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
	// errorMsg
	ServiceErrorMsg string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.List = m.List[:0]
	m.ServiceErrorCode = ""
	m.ServiceErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse)
	},
}

// GetTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse
func GetTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse() *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse {
	return poolTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse.Get().(*TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse)
}

// ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse 将 TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse(v *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse.Put(v)
}
