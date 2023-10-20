package tmallitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemsExtendSearchAPIResponse 搜索天猫商品 API返回值
// tmall.items.extend.search
//
// 提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
type TmallItemsExtendSearchAPIResponse struct {
	model.CommonResponse
	TmallItemsExtendSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemsExtendSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemsExtendSearchAPIResponseModel).Reset()
}

// TmallItemsExtendSearchAPIResponseModel is 搜索天猫商品 成功返回结果
type TmallItemsExtendSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_items_extend_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 品牌列表
	BrandList []TmallBrand `json:"brand_list,omitempty" xml:"brand_list>tmall_brand,omitempty"`
	// 商品列表
	ItemList []TmallExtendSearchItem `json:"item_list,omitempty" xml:"item_list>tmall_extend_search_item,omitempty"`
	// 类目列表
	CatList []TmallCat `json:"cat_list,omitempty" xml:"cat_list>tmall_cat,omitempty"`
	// 查询条件
	Q string `json:"q,omitempty" xml:"q,omitempty"`
	// 总商品数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemsExtendSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BrandList = m.BrandList[:0]
	m.ItemList = m.ItemList[:0]
	m.CatList = m.CatList[:0]
	m.Q = ""
	m.TotalResults = 0
}

var poolTmallItemsExtendSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemsExtendSearchAPIResponse)
	},
}

// GetTmallItemsExtendSearchAPIResponse 从 sync.Pool 获取 TmallItemsExtendSearchAPIResponse
func GetTmallItemsExtendSearchAPIResponse() *TmallItemsExtendSearchAPIResponse {
	return poolTmallItemsExtendSearchAPIResponse.Get().(*TmallItemsExtendSearchAPIResponse)
}

// ReleaseTmallItemsExtendSearchAPIResponse 将 TmallItemsExtendSearchAPIResponse 保存到 sync.Pool
func ReleaseTmallItemsExtendSearchAPIResponse(v *TmallItemsExtendSearchAPIResponse) {
	v.Reset()
	poolTmallItemsExtendSearchAPIResponse.Put(v)
}
