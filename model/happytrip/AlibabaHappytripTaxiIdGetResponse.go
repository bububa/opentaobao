package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取请求id APIResponse
alibaba.happytrip.taxi.id.get

获取订单号
*/
type AlibabaHappytripTaxiIdGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiIdGetResponse
}

type AlibabaHappytripTaxiIdGetResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_id_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // id获取结果
    
    Data   *GetIdResult `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
