package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretaildevicetradeshipAPIResponse 贩卖机掉货成功通知 API返回值
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
type AlibabaretaildevicetradeshipAPIResponse struct {
	model.CommonResponse
	AlibabaretaildevicetradeshipAPIResponseModel
}

// AlibabaretaildevicetradeshipAPIResponseModel is 贩卖机掉货成功通知 成功返回结果
type AlibabaretaildevicetradeshipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_trade_ship_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaretaildevicetradeshipResult `json:"result,omitempty" xml:"result,omitempty"`
}
