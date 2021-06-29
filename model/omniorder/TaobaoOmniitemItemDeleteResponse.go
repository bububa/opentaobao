package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品删除 APIResponse
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作
*/
type TaobaoOmniitemItemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemItemDeleteResponse
}

type TaobaoOmniitemItemDeleteResponse struct {
    XMLName xml.Name `xml:"omniitem_item_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemItemDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
