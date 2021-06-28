package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加购URL获取 APIResponse
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
type TaobaoItemCarturlGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemCarturlGetResponse
}

type TaobaoItemCarturlGetResponse struct {
    XMLName xml.Name `xml:"item_carturl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 加购的URL地址
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
