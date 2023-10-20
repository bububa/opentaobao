package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemCategoryGetAPIResponse 全渠道商品轻发布类目信息 API返回值
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
type TaobaoOmniitemCategoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniitemCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniitemCategoryGetAPIResponseModel).Reset()
}

// TaobaoOmniitemCategoryGetAPIResponseModel is 全渠道商品轻发布类目信息 成功返回结果
type TaobaoOmniitemCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemCategoryGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniitemCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniitemCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemCategoryGetAPIResponse)
	},
}

// GetTaobaoOmniitemCategoryGetAPIResponse 从 sync.Pool 获取 TaobaoOmniitemCategoryGetAPIResponse
func GetTaobaoOmniitemCategoryGetAPIResponse() *TaobaoOmniitemCategoryGetAPIResponse {
	return poolTaobaoOmniitemCategoryGetAPIResponse.Get().(*TaobaoOmniitemCategoryGetAPIResponse)
}

// ReleaseTaobaoOmniitemCategoryGetAPIResponse 将 TaobaoOmniitemCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniitemCategoryGetAPIResponse(v *TaobaoOmniitemCategoryGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniitemCategoryGetAPIResponse.Put(v)
}
