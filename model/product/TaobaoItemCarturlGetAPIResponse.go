package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加购URL获取 API返回值 
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
type TaobaoItemCarturlGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemCarturlGetAPIResponseModel
}

// 加购URL获取 成功返回结果
type TaobaoItemCarturlGetAPIResponseModel struct {
    XMLName xml.Name `xml:"item_carturl_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 加购的URL地址
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
