package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品全量更新 APIResponse
taobao.omniitem.item.fullupdate

全渠道商品全量更新，仅适用于全渠道门店商品
需要全量传入商品相关所有参数，更新时会根据传入的字段进行全量更新
对于SKU信息，会以skus对象进行判断，若传入的skus对象的sku为商品之前未包含的，则新增SKU，如果原先商品有该sku但现在没有传，则删除该SKU。所有传入的SKU信息要么全部均传入skuId，要么全部都不传入skuId。对于新增SKU的场景，目前无需传入SKUID，会根据传入的销售属性自动对应
*/
type TaobaoOmniitemItemFullupdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemItemFullupdateResponse
}

type TaobaoOmniitemItemFullupdateResponse struct {
    XMLName xml.Name `xml:"omniitem_item_fullupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoOmniitemItemFullupdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
