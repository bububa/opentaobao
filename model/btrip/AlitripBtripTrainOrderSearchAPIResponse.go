package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainOrderSearchAPIResponse 获取火车票订单列表 API返回值
// alitrip.btrip.train.order.search
//
// 第三方OA厂商获取自己的火车票数据
type AlitripBtripTrainOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripTrainOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripTrainOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripTrainOrderSearchAPIResponseModel).Reset()
}

// AlitripBtripTrainOrderSearchAPIResponseModel is 获取火车票订单列表 成功返回结果
type AlitripBtripTrainOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_train_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripTrainOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripTrainOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripTrainOrderSearchAPIResponse)
	},
}

// GetAlitripBtripTrainOrderSearchAPIResponse 从 sync.Pool 获取 AlitripBtripTrainOrderSearchAPIResponse
func GetAlitripBtripTrainOrderSearchAPIResponse() *AlitripBtripTrainOrderSearchAPIResponse {
	return poolAlitripBtripTrainOrderSearchAPIResponse.Get().(*AlitripBtripTrainOrderSearchAPIResponse)
}

// ReleaseAlitripBtripTrainOrderSearchAPIResponse 将 AlitripBtripTrainOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripTrainOrderSearchAPIResponse(v *AlitripBtripTrainOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripTrainOrderSearchAPIResponse.Put(v)
}
