package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuSortGetAPIResponse sku销售属性顺序获取 API返回值
// tmall.item.sku.sort.get
//
// sku销售属性顺序获取
type TmallItemSkuSortGetAPIResponse struct {
	model.CommonResponse
	TmallItemSkuSortGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSkuSortGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSkuSortGetAPIResponseModel).Reset()
}

// TmallItemSkuSortGetAPIResponseModel is sku销售属性顺序获取 成功返回结果
type TmallItemSkuSortGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_sort_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallItemSkuSortGetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSkuSortGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSkuSortGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSkuSortGetAPIResponse)
	},
}

// GetTmallItemSkuSortGetAPIResponse 从 sync.Pool 获取 TmallItemSkuSortGetAPIResponse
func GetTmallItemSkuSortGetAPIResponse() *TmallItemSkuSortGetAPIResponse {
	return poolTmallItemSkuSortGetAPIResponse.Get().(*TmallItemSkuSortGetAPIResponse)
}

// ReleaseTmallItemSkuSortGetAPIResponse 将 TmallItemSkuSortGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSkuSortGetAPIResponse(v *TmallItemSkuSortGetAPIResponse) {
	v.Reset()
	poolTmallItemSkuSortGetAPIResponse.Put(v)
}
