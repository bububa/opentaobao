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
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_conveyor_wcsbtoc_containerscannedbyconveyor_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ServiceErrorCode   string `json:"service_error_code,omitempty" xml:"