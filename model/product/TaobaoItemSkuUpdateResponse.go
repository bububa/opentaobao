package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新SKU信息 APIResponse
taobao.item.sku.update

*更新一个sku的数据 <br/>*需要更新的sku通过属性properties进行匹配查找 <br/>*商品的数量和价格必须大于等于0 <br/>*sku记录会更新到指定的num_iid对应的商品中 <br/>*num_iid对应的商品必须属于当前的会话用户
*/
type TaobaoItemSkuUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoItemSkuUpdateResponse
}

type TaobaoItemSkuUpdateResponse struct {
    XMLName xml.Name `xml:"item_sku_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品Sku
    
    Sku   *Sku `json:"sku,omitempty" xml:"sku,omitempty"`

    
}
