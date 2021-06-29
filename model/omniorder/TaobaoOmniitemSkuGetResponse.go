package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道门店商品sku APIResponse
taobao.omniitem.sku.get

通过skuId或者skuOutId查询全渠道门店商品sku信息
*/
type TaobaoOmniitemSkuGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemSkuGetResponse
}

type TaobaoOmniitemSkuGetResponse struct {
    XMLName xml.Name `xml:"omniitem_sku_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoOmniitemSkuGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
