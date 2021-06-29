package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查看订单详情 APIResponse
taobao.trade.drug.order.get

商家查看订单详情
*/
type TaobaoTradeDrugOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugOrderGetResponse
}

type TaobaoTradeDrugOrderGetResponse struct {
    XMLName xml.Name `xml:"trade_drug_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoTradeDrugOrderGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
