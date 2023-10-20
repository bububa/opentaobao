package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainCitySuggestAPIResponse 火车票城市搜索 API返回值
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
type AlitripBtripTrainCitySuggestAPIResponse struct {
	model.CommonResponse
	AlitripBtripTrainCitySuggestAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripTrainCitySuggestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripTrainCitySuggestAPIResponseModel).Reset()
}

// AlitripBtripTrainCitySuggestAPIResponseModel is 火车票城市搜索 成功返回结果
type AlitripBtripTrainCitySuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_train_city_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripTrainCitySuggestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripTrainCitySuggestAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripTrainCitySuggestAPIResponse)
	},
}

// GetAlitripBtripTrainCitySuggestAPIResponse 从 sync.Pool 获取 AlitripBtripTrainCitySuggestAPIResponse
func GetAlitripBtripTrainCitySuggestAPIResponse() *AlitripBtripTrainCitySuggestAPIResponse {
	return poolAlitripBtripTrainCitySuggestAPIResponse.Get().(*AlitripBtripTrainCitySuggestAPIResponse)
}

// ReleaseAlitripBtripTrainCitySuggestAPIResponse 将 AlitripBtripTrainCitySuggestAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripTrainCitySuggestAPIResponse(v *AlitripBtripTrainCitySuggestAPIResponse) {
	v.Reset()
	poolAlitripBtripTrainCitySuggestAPIResponse.Put(v)
}
