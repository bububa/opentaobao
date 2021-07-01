package tmallitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemsExtendSearchAPIResponse
搜索天猫商品 API返回值
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search */
type TmallItemsExtendSearchAPIResponse struct {
	model.CommonResponse
	TmallItemsExtendSearchAPIResponseModel
}

// TmallItemsExtendSearchAPIResponseModel is 搜索天猫商品 成功返回结果
type TmallItemsExtendSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_items_extend_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品列表
	ItemList []TmallExtendSearchItem `json:"item_list,omitempty" xml:"item_list>tmall_extend_search_item,omitempty"`
	// 总商品数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 查询条件
	Q string `json:"q,omitempty" xml:"q,omitempty"`
	// 类目列表
	CatList []TmallCat `json:"cat_list,omitempty" xml:"cat_list>tmall_cat,omitempty"`
	// 品牌列表
	BrandList []TmallBrand `json:"brand_list,omitempty" xml:"brand_list>tmall_brand,omitempty"`
}
