package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品解绑 APIResponse
taobao.opentrade.special.items.unbind

专属下单场景商品解绑
*/
type TaobaoOpentradeSpecialItemsUnbindAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeSpecialItemsUnbindResponse
}

type TaobaoOpentradeSpecialItemsUnbindResponse struct {
    XMLName xml.Name `xml:"opentrade_special_items_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 解绑返回结构
    
    Results   []ItemUnBindResult `json:"results,omitempty" xml:"results>item_un_bind_result,omitempty"`
    
    
}
