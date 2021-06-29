package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品轻发布类目信息 APIResponse
taobao.omniitem.category.get

全渠道商品轻发布类目信息
*/
type TaobaoOmniitemCategoryGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemCategoryGetResponse
}

type TaobaoOmniitemCategoryGetResponse struct {
    XMLName xml.Name `xml:"omniitem_category_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemCategoryGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
