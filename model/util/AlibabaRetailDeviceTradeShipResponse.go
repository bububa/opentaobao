package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机掉货成功通知 APIResponse
alibaba.retail.device.trade.ship

贩卖机发货
*/
type AlibabaRetailDeviceTradeShipAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailDeviceTradeShipResponse `json:"alibaba_retail_device_trade_ship_response,omitempty"`
}

type AlibabaRetailDeviceTradeShipResponse struct {

    // result
    Result   *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty"`

}
