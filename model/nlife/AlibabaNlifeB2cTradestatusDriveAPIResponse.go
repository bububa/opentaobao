package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradestatusDriveAPIResponse b2c订单状态驱动 API返回值
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
type AlibabaNlifeB2cTradestatusDriveAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradestatusDriveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradestatusDriveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cTradestatusDriveAPIResponseModel).Reset()
}

// AlibabaNlifeB2cTradestatusDriveAPIResponseModel is b2c订单状态驱动 成功返回结果
type AlibabaNlifeB2cTradestatusDriveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_tradestatus_drive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradestatusDriveAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaNlifeB2cTradestatusDriveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cTradestatusDriveAPIResponse)
	},
}

// GetAlibabaNlifeB2cTradestatusDriveAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cTradestatusDriveAPIResponse
func GetAlibabaNlifeB2cTradestatusDriveAPIResponse() *AlibabaNlifeB2cTradestatusDriveAPIResponse {
	return poolAlibabaNlifeB2cTradestatusDriveAPIResponse.Get().(*AlibabaNlifeB2cTradestatusDriveAPIResponse)
}

// ReleaseAlibabaNlifeB2cTradestatusDriveAPIResponse 将 AlibabaNlifeB2cTradestatusDriveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cTradestatusDriveAPIResponse(v *AlibabaNlifeB2cTradestatusDriveAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cTradestatusDriveAPIResponse.Put(v)
}
