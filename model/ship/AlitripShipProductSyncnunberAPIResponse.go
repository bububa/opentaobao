package ship

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripShipProductSyncnunberAPIResponse 船票班次变更回调 API返回值
// alitrip.ship.product.syncnunber
//
// 船票班次变更回调
type AlitripShipProductSyncnunberAPIResponse struct {
	model.CommonResponse
	AlitripShipProductSyncnunberAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripShipProductSyncnunberAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripShipProductSyncnunberAPIResponseModel).Reset()
}

// AlitripShipProductSyncnunberAPIResponseModel is 船票班次变更回调 成功返回结果
type AlitripShipProductSyncnunberAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ship_product_syncnunber_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripShipProductSyncnunberAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAlitripShipProductSyncnunberAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripShipProductSyncnunberAPIResponse)
	},
}

// GetAlitripShipProductSyncnunberAPIResponse 从 sync.Pool 获取 AlitripShipProductSyncnunberAPIResponse
func GetAlitripShipProductSyncnunberAPIResponse() *AlitripShipProductSyncnunberAPIResponse {
	return poolAlitripShipProductSyncnunberAPIResponse.Get().(*AlitripShipProductSyncnunberAPIResponse)
}

// ReleaseAlitripShipProductSyncnunberAPIResponse 将 AlitripShipProductSyncnunberAPIResponse 保存到 sync.Pool
func ReleaseAlitripShipProductSyncnunberAPIResponse(v *AlitripShipProductSyncnunberAPIResponse) {
	v.Reset()
	poolAlitripShipProductSyncnunberAPIResponse.Put(v)
}
