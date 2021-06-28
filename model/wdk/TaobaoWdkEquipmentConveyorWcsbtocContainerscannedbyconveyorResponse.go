package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
容器被悬挂链扫描 APIResponse
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse `json:"wdk_equipment_conveyor_wcsbtoc_containerscannedbyconveyor_response,omitempty"` 
    TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse struct {

    // errorCode
    
    ServiceErrorCode   string `json:"service_error_code,omitempty"`
    

    // errorMsg
    
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`
    

    // list
    
    List  struct {
        WcsContainerScannedByConveyorDto  []WcsContainerScannedByConveyorDto `json:"wcs_container_scanned_by_conveyor_dto,omitempty"`
    } `json:"list,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorResponse struct {

    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty"`

    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`

    // list
    List   []WcsContainerScannedByConveyorDto `json:"list,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
