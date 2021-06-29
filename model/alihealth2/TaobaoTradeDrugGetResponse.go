package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家未确认订单列表 APIResponse
taobao.trade.drug.get

可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
*/
type TaobaoTradeDrugGetAPIResponse struct {
    model.CommonResponse
    TaobaoTradeDrugGetResponse
}

type TaobaoTradeDrugGetResponse struct {
    XMLName xml.Name `xml:"trade_drug_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询到的订单列表对象
    
    Result   *TaobaoTradeDrugGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
