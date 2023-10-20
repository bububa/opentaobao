package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemCombineGetAPIResponse 组合商品获取接口 API返回值
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
type TmallItemCombineGetAPIResponse struct {
	model.CommonResponse
	TmallItemCombineGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemCombineGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemCombineGetAPIResponseModel).Reset()
}

// TmallItemCombineGetAPIResponseModel is 组合商品获取接口 成功返回结果
type TmallItemCombineGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_combine_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemCombineGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTmallItemCombineGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemCombineGetAPIResponse)
	},
}

// GetTmallItemCombineGetAPIResponse 从 sync.Pool 获取 TmallItemCombineGetAPIResponse
func GetTmallItemCombineGetAPIResponse() *TmallItemCombineGetAPIResponse {
	return poolTmallItemCombineGetAPIResponse.Get().(*TmallItemCombineGetAPIResponse)
}

// ReleaseTmallItemCombineGetAPIResponse 将 TmallItemCombineGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemCombineGetAPIResponse(v *TmallItemCombineGetAPIResponse) {
	v.Reset()
	poolTmallItemCombineGetAPIResponse.Put(v)
}
