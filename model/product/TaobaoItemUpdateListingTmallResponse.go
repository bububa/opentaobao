package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.listing天猫分流 APIResponse
taobao.item.update.listing.tmall

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateListingTmallAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_update_listing_tmall_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 上架后返回的商品信息：返回的结果就是:num_iid和modified
    
    Item   *Item `json:"item,omitempty" xml:"