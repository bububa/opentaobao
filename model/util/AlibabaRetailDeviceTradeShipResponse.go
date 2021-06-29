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
    AlibabaRetailDeviceTradeShipResponse
}

type AlibabaRetailDeviceTradeShipResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_trade_ship_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaRetailDeviceTradeShipResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
