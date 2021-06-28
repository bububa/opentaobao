package tmallitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索天猫商品 APIResponse
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
*/
type TmallItemsExtendSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_items_extend_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品列表
    
    ItemList   []TmallExtendSearchItem `json:"item_list,omitempty" xml:"