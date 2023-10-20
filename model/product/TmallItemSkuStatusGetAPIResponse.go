package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSkuStatusGetAPIResponse 商品sku上下架查询 API返回值
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
type TmallItemSkuStatusGetAPIResponse struct {
	model.CommonResponse
	TmallItemSkuStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSkuStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSkuStatusGetAPIResponseModel).Reset()
}

// TmallItemSkuStatusGetAPIResponseModel is 商品sku上下架查询 成功返回结果
type TmallItemSkuStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sku_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TmallItemSkuStatusGetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSkuStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSkuStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSkuStatusGetAPIResponse)
	},
}

// GetTmallItemSkuStatusGetAPIResponse 从 sync.Pool 获取 TmallItemSkuStatusGetAPIResponse
func GetTmallItemSkuStatusGetAPIResponse() *TmallItemSkuStatusGetAPIResponse {
	return poolTmallItemSkuStatusGetAPIResponse.Get().(*TmallItemSkuStatusGetAPIResponse)
}

// ReleaseTmallItemSkuStatusGetAPIResponse 将 TmallItemSkuStatusGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSkuStatusGetAPIResponse(v *TmallItemSkuStatusGetAPIResponse) {
	v.Reset()
	poolTmallItemSkuStatusGetAPIResponse.Put(v)
}
