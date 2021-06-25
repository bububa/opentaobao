package tmallitem

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索天猫商品 APIResponse
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
*/
type TmallItemsExtendSearchAPIResponse struct {
    model.CommonResponse
    Response *TmallItemsExtendSearchResponse `json:"tmall_items_extend_search_response,omitempty"`
}

type TmallItemsExtendSearchResponse struct {

    // 商品列表
    ItemList   []TmallExtendSearchItem `json:"item_list,omitempty"`

    // 总商品数量
    TotalResults   int64 `json:"total_results,omitempty"`

    // 查询条件
    Q   string `json:"q,omitempty"`

    // 类目列表
    CatList   []TmallCat `json:"cat_list,omitempty"`

    // 品牌列表
    BrandList   []TmallBrand `json:"brand_list,omitempty"`

}
