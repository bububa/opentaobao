package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceTradeShipAPIResponse 贩卖机掉货成功通知 API返回值
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
type AlibabaRetailDeviceTradeShipAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceTradeShipAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceTradeShipAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailDeviceTradeShipAPIResponseModel).Reset()
}

// AlibabaRetailDeviceTradeShipAPIResponseModel is 贩卖机掉货成功通知 成功返回结果
type AlibabaRetailDeviceTradeShipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_trade_ship_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceTradeShipAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailDeviceTradeShipAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceTradeShipAPIResponse)
	},
}

// GetAlibabaRetailDeviceTradeShipAPIResponse 从 sync.Pool 获取 AlibabaRetailDeviceTradeShipAPIResponse
func GetAlibabaRetailDeviceTradeShipAPIResponse() *AlibabaRetailDeviceTradeShipAPIResponse {
	return poolAlibabaRetailDeviceTradeShipAPIResponse.Get().(*AlibabaRetailDeviceTradeShipAPIResponse)
}

// ReleaseAlibabaRetailDeviceTradeShipAPIResponse 将 AlibabaRetailDeviceTradeShipAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailDeviceTradeShipAPIResponse(v *AlibabaRetailDeviceTradeShipAPIResponse) {
	v.Reset()
	poolAlibabaRetailDeviceTradeShipAPIResponse.Put(v)
}
