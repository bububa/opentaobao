package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse 【商旅】火车票订单详情查询 API返回值
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
type AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponseModel).Reset()
}

// AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponseModel is 【商旅】火车票订单详情查询 成功返回结果
type AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_detail_search_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse)
	},
}

// GetAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse
func GetAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse() *AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse {
	return poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse.Get().(*AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse)
}

// ReleaseAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse 将 AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse(v *AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse.Put(v)
}
