package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设置设备标准属性状态 APIResponse
alibaba.alink.device.unify.status.set

操作用户绑定的设备
*/
type AlibabaAlinkDeviceUnifyStatusSetAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkDeviceUnifyStatusSetResponse
}

type AlibabaAlinkDeviceUnifyStatusSetResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_device_unify_status_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
