package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设置设备标准属性状态 APIResponse
alibaba.alink.device.unify.status.set

操作用户绑定的设备
*/
type AlibabaAlinkDeviceUnifyStatusSetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkDeviceUnifyStatusSetResponse `json:"alibaba_alink_device_unify_status_set_response,omitempty"` 
    AlibabaAlinkDeviceUnifyStatusSetResponse
}

/* model for simplify = false
type AlibabaAlinkDeviceUnifyStatusSetResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkDeviceUnifyStatusSetResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
