package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiCityGetidAPIResponse 获取城市id API返回值
// alibaba.happytrip.taxi.city.getid
//
// 通过经纬度坐标返回城市id
type AlibabaHappytripTaxiCityGetidAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiCityGetidAPIResponseModel
}

// AlibabaHappytripTaxiCityGetidAPIResponseModel is 获取城市id 成功返回结果
type AlibabaHappytripTaxiCityGetidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_city_getid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 结果对象
	Data *CityGetIdResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}
