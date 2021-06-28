package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机掉货成功通知 APIResponse
alibaba.retail.device.trade.ship

贩卖机发货
*/
type AlibabaRetailDeviceTradeShipAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_device_trade_ship_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty" xml:"