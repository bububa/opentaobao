package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainTrainTradeAPIResponse 商旅火车票交易流水接口 API返回值
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
type AlitripBtripOpenSupplychainTrainTradeAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenSupplychainTrainTradeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainTrainTradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenSupplychainTrainTradeAPIResponseModel).Reset()
}

// AlitripBtripOpenSupplychainTrainTradeAPIResponseModel is 商旅火车票交易流水接口 成功返回结果
type AlitripBtripOpenSupplychainTrainTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_train_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenSupplychainTrainTradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripOpenSupplychainTrainTradeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenSupplychainTrainTradeAPIResponse)
	},
}

// GetAlitripBtripOpenSupplychainTrainTradeAPIResponse 从 sync.Pool 获取 AlitripBtripOpenSupplychainTrainTradeAPIResponse
func GetAlitripBtripOpenSupplychainTrainTradeAPIResponse() *AlitripBtripOpenSupplychainTrainTradeAPIResponse {
	return poolAlitripBtripOpenSupplychainTrainTradeAPIResponse.Get().(*AlitripBtripOpenSupplychainTrainTradeAPIResponse)
}

// ReleaseAlitripBtripOpenSupplychainTrainTradeAPIResponse 将 AlitripBtripOpenSupplychainTrainTradeAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenSupplychainTrainTradeAPIResponse(v *AlitripBtripOpenSupplychainTrainTradeAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenSupplychainTrainTradeAPIResponse.Put(v)
}
