package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainCityAPIResponse 火车站数据查询 API返回值
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
type AlitripBtripSupplychainTrainCityAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainTrainCityAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainCityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainTrainCityAPIResponseModel).Reset()
}

// AlitripBtripSupplychainTrainCityAPIResponseModel is 火车站数据查询 成功返回结果
type AlitripBtripSupplychainTrainCityAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_city_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Result *OpenApiSuggestRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果信息
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainCityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripSupplychainTrainCityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainTrainCityAPIResponse)
	},
}

// GetAlitripBtripSupplychainTrainCityAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainTrainCityAPIResponse
func GetAlitripBtripSupplychainTrainCityAPIResponse() *AlitripBtripSupplychainTrainCityAPIResponse {
	return poolAlitripBtripSupplychainTrainCityAPIResponse.Get().(*AlitripBtripSupplychainTrainCityAPIResponse)
}

// ReleaseAlitripBtripSupplychainTrainCityAPIResponse 将 AlitripBtripSupplychainTrainCityAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainTrainCityAPIResponse(v *AlitripBtripSupplychainTrainCityAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainTrainCityAPIResponse.Put(v)
}
