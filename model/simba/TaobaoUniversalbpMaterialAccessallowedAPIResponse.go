package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialAccessallowedAPIResponse 物料准入判断 API返回值
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
type TaobaoUniversalbpMaterialAccessallowedAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpMaterialAccessallowedAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMaterialAccessallowedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpMaterialAccessallowedAPIResponseModel).Reset()
}

// TaobaoUniversalbpMaterialAccessallowedAPIResponseModel is 物料准入判断 成功返回结果
type TaobaoUniversalbpMaterialAccessallowedAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_material_accessallowed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpMaterialAccessallowedTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMaterialAccessallowedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpMaterialAccessallowedAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMaterialAccessallowedAPIResponse)
	},
}

// GetTaobaoUniversalbpMaterialAccessallowedAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpMaterialAccessallowedAPIResponse
func GetTaobaoUniversalbpMaterialAccessallowedAPIResponse() *TaobaoUniversalbpMaterialAccessallowedAPIResponse {
	return poolTaobaoUniversalbpMaterialAccessallowedAPIResponse.Get().(*TaobaoUniversalbpMaterialAccessallowedAPIResponse)
}

// ReleaseTaobaoUniversalbpMaterialAccessallowedAPIResponse 将 TaobaoUniversalbpMaterialAccessallowedAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpMaterialAccessallowedAPIResponse(v *TaobaoUniversalbpMaterialAccessallowedAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpMaterialAccessallowedAPIResponse.Put(v)
}
