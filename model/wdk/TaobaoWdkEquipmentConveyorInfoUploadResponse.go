package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口仓库悬挂链信息上报 APIResponse
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentConveyorInfoUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentConveyorInfoUploadResponse `json:"wdk_equipment_conveyor_info_upload_response,omitempty"` 
    TaobaoWdkEquipmentConveyorInfoUploadResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentConveyorInfoUploadResponse struct {

    // model
    
    Model   string `json:"model,omitempty"`
    

    // errorCode
    
    ServiceErrorCode   string `json:"service_error_code,omitempty"`
    

    // errorMsg
    
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoWdkEquipmentConveyorInfoUploadResponse struct {

    // model
    Model   string `json:"model,omitempty"`

    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty"`

    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
