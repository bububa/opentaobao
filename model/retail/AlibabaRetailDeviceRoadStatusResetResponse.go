package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机货道解锁 APIResponse
alibaba.retail.device.road.status.reset

贩卖机货道解锁
*/
type AlibabaRetailDeviceRoadStatusResetAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDeviceRoadStatusResetResponse
}

type AlibabaRetailDeviceRoadStatusResetResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_road_status_reset_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
}
