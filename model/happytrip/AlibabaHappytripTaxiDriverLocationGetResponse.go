package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机位置 APIResponse
alibaba.happytrip.taxi.driver.location.get

获取司机实时位置
*/
type AlibabaHappytripTaxiDriverLocationGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiDriverLocationGetResponse
}

type AlibabaHappytripTaxiDriverLocationGetResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_location_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 司机位置
    
    Data   *AlibabaHappytripTaxiDriverLocationGetStruct `json:"data,omitempty" xml:"data,omitempty"`

    
}
