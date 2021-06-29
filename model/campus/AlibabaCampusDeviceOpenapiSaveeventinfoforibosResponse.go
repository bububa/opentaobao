package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
saveeventinfoforibos APIResponse
alibaba.campus.device.openapi.saveeventinfoforibos

IBos的事件信息上报与反馈的处理接口
*/
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceOpenapiSaveeventinfoforibosResponse
}

type AlibabaCampusDeviceOpenapiSaveeventinfoforibosResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_openapi_saveeventinfoforibos_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // errorMsg
    
    RequestErrorMsg   string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`

    
    // success
    
    RequestSuccess   bool `json:"request_success,omitempty" xml:"request_success,omitempty"`

    
}
