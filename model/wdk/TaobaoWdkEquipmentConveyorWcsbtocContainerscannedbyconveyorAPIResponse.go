package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse
容器被悬挂链扫描 API返回值
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描 */
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel
}

// TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel is 容器被悬挂链扫描 成功返回结果
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_conveyor_wcsbtoc_containerscannedbyconveyor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ServiceErrorCode string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
	// errorMsg
	ServiceErrorMsg string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
	// list
	List []WcsContainerScannedByConveyorDto `json:"list,omitempty" xml:"list>wcs_container_scanned_by_conveyor_dto,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
