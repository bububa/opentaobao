package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市id APIResponse
alibaba.happytrip.taxi.city.getid

通过经纬度坐标返回城市id
*/
type AlibabaHappytripTaxiCityGetidAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiCityGetidResponse
}

type AlibabaHappytripTaxiCityGetidResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_city_getid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Data   *CityGetIdResult `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
