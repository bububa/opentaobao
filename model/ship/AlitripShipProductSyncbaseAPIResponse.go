package ship

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripShipProductSyncbaseAPIResponse 基础信息修改回调 API返回值
// alitrip.ship.product.syncbase
//
// 基础信息修改回调
type AlitripShipProductSyncbaseAPIResponse struct {
	model.CommonResponse
	AlitripShipProductSyncbaseAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripShipProductSyncbaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripShipProductSyncbaseAPIResponseModel).Reset()
}

// AlitripShipProductSyncbaseAPIResponseModel is 基础信息修改回调 成功返回结果
type AlitripShipProductSyncbaseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ship_product_syncbase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripShipProductSyncbaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAlitripShipProductSyncbaseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripShipProductSyncbaseAPIResponse)
	},
}

// GetAlitripShipProductSyncbaseAPIResponse 从 sync.Pool 获取 AlitripShipProductSyncbaseAPIResponse
func GetAlitripShipProductSyncbaseAPIResponse() *AlitripShipProductSyncbaseAPIResponse {
	return poolAlitripShipProductSyncbaseAPIResponse.Get().(*AlitripShipProductSyncbaseAPIResponse)
}

// ReleaseAlitripShipProductSyncbaseAPIResponse 将 AlitripShipProductSyncbaseAPIResponse 保存到 sync.Pool
func ReleaseAlitripShipProductSyncbaseAPIResponse(v *AlitripShipProductSyncbaseAPIResponse) {
	v.Reset()
	poolAlitripShipProductSyncbaseAPIResponse.Put(v)
}
