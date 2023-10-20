package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuStatusUpdateAPIResponse 商品sku状态更新 API返回值
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
type TmallItemSkuStatusUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSkuStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSkuStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSkuStatusUpdateAPIResponseModel).Reset()
}

// TmallItemSkuStatusUpdateAPIResponseModel is 商品sku状态更新 成功返回结果
type TmallItemSkuStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallItemSkuStatusUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSkuStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSkuStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSkuStatusUpdateAPIResponse)
	},
}

// GetTmallItemSkuStatusUpdateAPIResponse 从 sync.Pool 获取 TmallItemSkuStatusUpdateAPIResponse
func GetTmallItemSkuStatusUpdateAPIResponse() *TmallItemSkuStatusUpdateAPIResponse {
	return poolTmallItemSkuStatusUpdateAPIResponse.Get().(*TmallItemSkuStatusUpdateAPIResponse)
}

// ReleaseTmallItemSkuStatusUpdateAPIResponse 将 TmallItemSkuStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSkuStatusUpdateAPIResponse(v *TmallItemSkuStatusUpdateAPIResponse) {
	v.Reset()
	poolTmallItemSkuStatusUpdateAPIResponse.Put(v)
}
