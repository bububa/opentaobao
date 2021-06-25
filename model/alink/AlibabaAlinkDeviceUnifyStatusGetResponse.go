package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询设备标准属性最新状态 APIResponse
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态
*/
type AlibabaAlinkDeviceUnifyStatusGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlinkDeviceUnifyStatusGetResponse `json:"alibaba_alink_device_unify_status_get_response,omitempty"`
}

type AlibabaAlinkDeviceUnifyStatusGetResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
