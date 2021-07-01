package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机位置 API返回值 
alibaba.happytrip.taxi.driver.location.get

获取司机实时位置
*/
type AlibabaHappytripTaxiDriverLocationGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiDriverLocationGetAPIResponseModel
}

// 司机位置 成功返回结果
type AlibabaHappytripTaxiDriverLocationGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_location_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`
    // 错误信息
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
    // 司机位置
    Data   *AlibabaHappytripTaxiDriverLocationGetStruct `json:"data,omitempty" xml:"data,omitempty"`
}
