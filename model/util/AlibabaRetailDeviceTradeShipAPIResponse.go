package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceTradeShipAPIResponse
贩卖机掉货成功通知 API返回值
alibaba.retail.device.trade.ship

贩卖机发货 */
type AlibabaRetailDeviceTradeShipAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceTradeShipAPIResponseModel
}

// AlibabaRetailDeviceTradeShipAPIResponseModel is 贩卖机掉货成功通知 成功返回结果
type AlibabaRetailDeviceTradeShipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_trade_ship_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty" xml:"result,omitempty"`
}
