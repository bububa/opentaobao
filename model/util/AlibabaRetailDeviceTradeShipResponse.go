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
    // Response *AlibabaRetailDeviceTradeShipResponse `json:"alibaba_retail_device_trade_ship_response,omitempty"` 
    AlibabaRetailDeviceTradeShipResponse
}

/* model for simplify = false
type AlibabaRetailDeviceTradeShipResponse struct {

    // result
    
    Result  *struct {
        AlibabaRetailDeviceTradeShipResult  *AlibabaRetailDeviceTradeShipResult `json:"alibaba_retail_device_trade_ship_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailDeviceTradeShipResponse struct {

    // result
    Result   *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty"`

}
