package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTransferOrderDetailAPIResponse 接送订单详情接口 API返回值
// alitrip.transfer.order.detail
//
// 接送订单详情接口
type AlitripTransferOrderDetailAPIResponse struct {
	model.CommonResponse
	AlitripTransferOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTransferOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTransferOrderDetailAPIResponseModel).Reset()
}

// AlitripTransferOrderDetailAPIResponseModel is 接送订单详情接口 成功返回结果
type AlitripTransferOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_transfer_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *TransferOrderDetailCallbackRsp `json:"data,omitempty" xml:"data,omitempty"`
	// 信息code
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTransferOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Data = nil
	m.MessageCode = 0
}

var poolAlitripTransferOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTransferOrderDetailAPIResponse)
	},
}

// GetAlitripTransferOrderDetailAPIResponse 从 sync.Pool 获取 AlitripTransferOrderDetailAPIResponse
func GetAlitripTransferOrderDetailAPIResponse() *AlitripTransferOrderDetailAPIResponse {
	return poolAlitripTransferOrderDetailAPIResponse.Get().(*AlitripTransferOrderDetailAPIResponse)
}

// ReleaseAlitripTransferOrderDetailAPIResponse 将 AlitripTransferOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripTransferOrderDetailAPIResponse(v *AlitripTransferOrderDetailAPIResponse) {
	v.Reset()
	poolAlitripTransferOrderDetailAPIResponse.Put(v)
}
