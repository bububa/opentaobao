package category

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemCatpropsModificationGetAPIResponse 查询商品类目属性变更 API返回值
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
type TaobaoItemCatpropsModificationGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemCatpropsModificationGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemCatpropsModificationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemCatpropsModificationGetAPIResponseModel).Reset()
}

// TaobaoItemCatpropsModificationGetAPIResponseModel is 查询商品类目属性变更 成功返回结果
type TaobaoItemCatpropsModificationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_catprops_modification_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Results []PropsModificationResult `json:"results,omitempty" xml:"results>props_modification_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemCatpropsModificationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoItemCatpropsModificationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemCatpropsModificationGetAPIResponse)
	},
}

// GetTaobaoItemCatpropsModificationGetAPIResponse 从 sync.Pool 获取 TaobaoItemCatpropsModificationGetAPIResponse
func GetTaobaoItemCatpropsModificationGetAPIResponse() *TaobaoItemCatpropsModificationGetAPIResponse {
	return poolTaobaoItemCatpropsModificationGetAPIResponse.Get().(*TaobaoItemCatpropsModificationGetAPIResponse)
}

// ReleaseTaobaoItemCatpropsModificationGetAPIResponse 将 TaobaoItemCatpropsModificationGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemCatpropsModificationGetAPIResponse(v *TaobaoItemCatpropsModificationGetAPIResponse) {
	v.Reset()
	poolTaobaoItemCatpropsModificationGetAPIResponse.Put(v)
}
