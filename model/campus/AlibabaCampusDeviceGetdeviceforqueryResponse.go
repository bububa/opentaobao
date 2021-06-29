package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下发设备的分页接口(无需AOP控制) APIResponse
alibaba.campus.device.getdeviceforquery

下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制)
*/
type AlibabaCampusDeviceGetdeviceforqueryAPIResponse struct {
    model.CommonResponse
    AlibabaCampusDeviceGetdeviceforqueryResponse
}

type AlibabaCampusDeviceGetdeviceforqueryResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_device_getdeviceforquery_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
