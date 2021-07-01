package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口仓库悬挂链信息上报 API返回值 
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传
*/
type TaobaoWdkEquipmentConveyorInfoUploadAPIResponse struct {
    model.CommonResponse
    TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel
}

// 五道口仓库悬挂链信息上报 成功返回结果
type TaobaoWdkEquipmentConveyorInfoUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"wdk_equipment_conveyor_info_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // model
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
