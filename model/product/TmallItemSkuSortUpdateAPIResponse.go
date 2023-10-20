package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuSortUpdateAPIResponse 商品销售属性排序更新 API返回值
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
type TmallItemSkuSortUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSkuSortUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSkuSortUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSkuSortUpdateAPIResponseModel).Reset()
}

// TmallItemSkuSortUpdateAPIResponseModel is 商品销售属性排序更新 成功返回结果
type TmallItemSkuSortUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_sort_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallItemSkuSortUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSkuSortUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSkuSortUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSkuSortUpdateAPIResponse)
	},
}

// GetTmallItemSkuSortUpdateAPIResponse 从 sync.Pool 获取 TmallItemSkuSortUpdateAPIResponse
func GetTmallItemSkuSortUpdateAPIResponse() *TmallItemSkuSortUpdateAPIResponse {
	return poolTmallItemSkuSortUpdateAPIResponse.Get().(*TmallItemSkuSortUpdateAPIResponse)
}

// ReleaseTmallItemSkuSortUpdateAPIResponse 将 TmallItemSkuSortUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSkuSortUpdateAPIResponse(v *TmallItemSkuSortUpdateAPIResponse) {
	v.Reset()
	poolTmallItemSkuSortUpdateAPIResponse.Put(v)
}
