package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备标准属性最新状态 APIResponse
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态
*/
type AlibabaAlinkDeviceUnifyStatusGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alink_device_unify_status_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"