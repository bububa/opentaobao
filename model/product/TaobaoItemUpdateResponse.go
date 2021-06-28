package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品信息 APIResponse
taobao.item.update

根据传入的num_iid更新对应的商品的数据。
传入的num_iid所对应的商品必须属于当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。
*/
type TaobaoItemUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoItemUpdateResponse
}

type TaobaoItemUpdateResponse struct {
    XMLName xml.Name `xml:"item_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品结构
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`

    
}
