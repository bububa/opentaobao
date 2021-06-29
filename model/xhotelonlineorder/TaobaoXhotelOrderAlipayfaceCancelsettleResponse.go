package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单取消结算接口 APIResponse
taobao.xhotel.order.alipayface.cancelsettle

信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
*/
type TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceCancelsettleResponse
}

type TaobaoXhotelOrderAlipayfaceCancelsettleResponse struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_cancelsettle_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
