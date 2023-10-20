package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverLocationGetAPIResponse 司机位置 API返回值
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
type AlibabaHappytripTaxiDriverLocationGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiDriverLocationGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverLocationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiDriverLocationGetAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiDriverLocationGetAPIResponseModel is 司机位置 成功返回结果
type AlibabaHappytripTaxiDriverLocationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_location_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 司机位置
	Data *AlibabaHappytripTaxiDriverLocationGetStruct `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverLocationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiDriverLocationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiDriverLocationGetAPIResponse)
	},
}

// GetAlibabaHappytripTaxiDriverLocationGetAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiDriverLocationGetAPIResponse
func GetAlibabaHappytripTaxiDriverLocationGetAPIResponse() *AlibabaHappytripTaxiDriverLocationGetAPIResponse {
	return poolAlibabaHappytripTaxiDriverLocationGetAPIResponse.Get().(*AlibabaHappytripTaxiDriverLocationGetAPIResponse)
}

// ReleaseAlibabaHappytripTaxiDriverLocationGetAPIResponse 将 AlibabaHappytripTaxiDriverLocationGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverLocationGetAPIResponse(v *AlibabaHappytripTaxiDriverLocationGetAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverLocationGetAPIResponse.Put(v)
}
