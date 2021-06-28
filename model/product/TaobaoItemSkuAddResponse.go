package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加SKU APIResponse
taobao.item.sku.add

新增一个sku到num_iid指定的商品中 <br/>传入的iid所对应的商品必须属于当前会话的用户
*/
type TaobaoItemSkuAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_sku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // sku
    
    Sku   *Sku `json:"sku,omitempty" xml:"