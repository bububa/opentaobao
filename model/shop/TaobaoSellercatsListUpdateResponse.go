package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新卖家自定义类目 APIResponse
taobao.sellercats.list.update

此API更新卖家店铺内自定义类目 <br/>注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
*/
type TaobaoSellercatsListUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"sellercats_list_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回sellercat数据结构中的：cid,modified
    
    SellerCat   *SellerCat `json:"seller_cat,omitempty" xml:"