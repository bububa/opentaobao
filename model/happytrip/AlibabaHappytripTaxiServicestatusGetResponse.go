package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商服务开通状态 APIResponse
alibaba.happytrip.taxi.servicestatus.get

获取服务供应商在每个地区的服务开通状态、支持的车型等
*/
type AlibabaHappytripTaxiServicestatusGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiServicestatusGetResponse
}

type AlibabaHappytripTaxiServicestatusGetResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_servicestatus_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 供应商服务状态数据
    
    Data   *ServiceStatusModel `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误代码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误消息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
