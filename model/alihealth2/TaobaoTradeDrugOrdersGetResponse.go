package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康获取某一药店全部订单 APIResponse
taobao.trade.drug.orders.get

阿里健康获取某一药店全部订单
*/
type TaobaoTradeDrugOrdersGetAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugOrdersGetResponse
}

type TaobaoTradeDrugOrdersGetResponse struct {
    XMLName xml.Name `xml:"trade_drug_orders_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询到的订单列表对象
    
    Result   *TaobaoTradeDrugOrdersGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
