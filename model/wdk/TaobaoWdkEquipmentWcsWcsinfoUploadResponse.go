package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链业务信息上传 APIResponse
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkEquipmentWcsWcsinfoUploadResponse `json:"wdk_equipment_wcs_wcsinfo_upload_response,omitempty"` 
    TaobaoWdkEquipmentWcsWcsinfoUploadResponse
}

/* model for simplify = false
type TaobaoWdkEquipmentWcsWcsinfoUploadResponse struct {

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

type TaobaoWdkEquipmentWcsWcsinfoUploadResponse struct {

    // model
    Model   string `json:"model,omitempty"`

    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty"`

    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
