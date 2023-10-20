package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiIdGetAPIResponse 获取请求id API返回值
// alibaba.happytrip.taxi.id.get
//
// 获取订单号
type AlibabaHappytripTaxiIdGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiIdGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiIdGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiIdGetAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiIdGetAPIResponseModel is 获取请求id 成功返回结果
type AlibabaHappytripTaxiIdGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_id_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// id获取结果
	Data *GetIdResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiIdGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiIdGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiIdGetAPIResponse)
	},
}

// GetAlibabaHappytripTaxiIdGetAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiIdGetAPIResponse
func GetAlibabaHappytripTaxiIdGetAPIResponse() *AlibabaHappytripTaxiIdGetAPIResponse {
	return poolAlibabaHappytripTaxiIdGetAPIResponse.Get().(*AlibabaHappytripTaxiIdGetAPIResponse)
}

// ReleaseAlibabaHappytripTaxiIdGetAPIResponse 将 AlibabaHappytripTaxiIdGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiIdGetAPIResponse(v *AlibabaHappytripTaxiIdGetAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiIdGetAPIResponse.Put(v)
}
