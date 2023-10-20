package ship

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripShipReturnNotifyAPIResponse 船票退票退款回填接口 API返回值
// alitrip.ship.return.notify
//
// 此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
type AlitripShipReturnNotifyAPIResponse struct {
	model.CommonResponse
	AlitripShipReturnNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripShipReturnNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripShipReturnNotifyAPIResponseModel).Reset()
}

// AlitripShipReturnNotifyAPIResponseModel is 船票退票退款回填接口 成功返回结果
type AlitripShipReturnNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ship_return_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 错误描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 结果
	RetSuccess bool `json:"ret_success,omitempty" xml:"ret_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripShipReturnNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RetSuccess = false
}

var poolAlitripShipReturnNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripShipReturnNotifyAPIResponse)
	},
}

// GetAlitripShipReturnNotifyAPIResponse 从 sync.Pool 获取 AlitripShipReturnNotifyAPIResponse
func GetAlitripShipReturnNotifyAPIResponse() *AlitripShipReturnNotifyAPIResponse {
	return poolAlitripShipReturnNotifyAPIResponse.Get().(*AlitripShipReturnNotifyAPIResponse)
}

// ReleaseAlitripShipReturnNotifyAPIResponse 将 AlitripShipReturnNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripShipReturnNotifyAPIResponse(v *AlitripShipReturnNotifyAPIResponse) {
	v.Reset()
	poolAlitripShipReturnNotifyAPIResponse.Put(v)
}
