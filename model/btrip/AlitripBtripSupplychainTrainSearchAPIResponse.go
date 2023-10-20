package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainSearchAPIResponse 【商旅】火车票订单查询 API返回值
// alitrip.btrip.supplychain.train.search
//
// 【商旅】火车票订单查询
type AlitripBtripSupplychainTrainSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainTrainSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainTrainSearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainTrainSearchAPIResponseModel is 【商旅】火车票订单查询 成功返回结果
type AlitripBtripSupplychainTrainSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口对外数据透出
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainTrainSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainTrainSearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainTrainSearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainTrainSearchAPIResponse
func GetAlitripBtripSupplychainTrainSearchAPIResponse() *AlitripBtripSupplychainTrainSearchAPIResponse {
	return poolAlitripBtripSupplychainTrainSearchAPIResponse.Get().(*AlitripBtripSupplychainTrainSearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainTrainSearchAPIResponse 将 AlitripBtripSupplychainTrainSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainTrainSearchAPIResponse(v *AlitripBtripSupplychainTrainSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainTrainSearchAPIResponse.Put(v)
}
