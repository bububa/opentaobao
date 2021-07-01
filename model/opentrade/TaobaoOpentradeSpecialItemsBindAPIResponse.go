package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品绑定 API返回值 
taobao.opentrade.special.items.bind

专属下单场景商品绑定
*/
type TaobaoOpentradeSpecialItemsBindAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeSpecialItemsBindAPIResponseModel
}

// 专属下单场景商品绑定 成功返回结果
type TaobaoOpentradeSpecialItemsBindAPIResponseModel struct {
    XMLName xml.Name `xml:"opentrade_special_items_bind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 绑定返回结构
    Results   []ItemBindResult `json:"results,omitempty" xml:"results>item_bind_result,omitempty"`
}
