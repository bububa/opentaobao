package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备信息获取 APIResponse
alibaba.retail.device.info.get

贩卖机设备信息获取
*/
type AlibabaRetailDeviceInfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDeviceInfoGetResponse
}

type AlibabaRetailDeviceInfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
}
