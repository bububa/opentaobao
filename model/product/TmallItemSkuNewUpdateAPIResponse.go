package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuNewUpdateAPIResponse 更新sku销售属性标新状态 API返回值
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
type TmallItemSkuNewUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSkuNewUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSkuNewUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSkuNewUpdateAPIResponseModel).Reset()
}

// TmallItemSkuNewUpdateAPIResponseModel is 更新sku销售属性标新状态 成功返回结果
type TmallItemSkuNewUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_new_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallItemSkuNewUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSkuNewUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSkuNewUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSkuNewUpdateAPIResponse)
	},
}

// GetTmallItemSkuNewUpdateAPIResponse 从 sync.Pool 获取 TmallItemSkuNewUpdateAPIResponse
func GetTmallItemSkuNewUpdateAPIResponse() *TmallItemSkuNewUpdateAPIResponse {
	return poolTmallItemSkuNewUpdateAPIResponse.Get().(*TmallItemSkuNewUpdateAPIResponse)
}

// ReleaseTmallItemSkuNewUpdateAPIResponse 将 TmallItemSkuNewUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSkuNewUpdateAPIResponse(v *TmallItemSkuNewUpdateAPIResponse) {
	v.Reset()
	poolTmallItemSkuNewUpdateAPIResponse.Put(v)
}
