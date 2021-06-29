package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商订单价格修改接口 APIResponse
alibaba.idle.isv.order.adjustprice

闲鱼用户通过授权的服务商修改订单价格和邮费
*/
type AlibabaIdleIsvOrderAdjustpriceAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvOrderAdjustpriceResponse
}

type AlibabaIdleIsvOrderAdjustpriceResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_order_adjustprice_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果
    
    Result   *AlibabaIdleIsvOrderAdjustpriceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
