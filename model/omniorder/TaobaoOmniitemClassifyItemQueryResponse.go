package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据分类查商品信息 APIResponse
taobao.omniitem.classify.item.query

商家根据分类查商品
*/
type TaobaoOmniitemClassifyItemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyItemQueryResponse
}

type TaobaoOmniitemClassifyItemQueryResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
