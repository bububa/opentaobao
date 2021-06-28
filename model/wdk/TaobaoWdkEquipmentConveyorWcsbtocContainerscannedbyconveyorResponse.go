package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被悬挂链扫描 APIResponse
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse
}

type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_wcsbtoc_containerscannedbyconveyor_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ServiceErrorCode   string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`

    
    // errorMsg
    
    ServiceErrorMsg   string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`

    
    // list
    
    List   []WcsContainerScannedByConveyorDto `json:"list,omitempty" xml:"list>wcs_container_scanned_by_conveyor_dto,omitempty"`
    
    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
