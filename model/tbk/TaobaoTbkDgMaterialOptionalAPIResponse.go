package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgMaterialOptionalAPIResponse 淘宝客-推广者-物料搜索 API返回值
// taobao.tbk.dg.material.optional
//
// 通用物料搜索API（导购）
type TaobaoTbkDgMaterialOptionalAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgMaterialOptionalAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgMaterialOptionalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgMaterialOptionalAPIResponseModel).Reset()
}

// TaobaoTbkDgMaterialOptionalAPIResponseModel is 淘宝客-推广者-物料搜索 成功返回结果
type TaobaoTbkDgMaterialOptionalAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_material_optional_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaoTbkDgMaterialOptionalMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_dg_material_optional_map_data,omitempty"`
	// 本地化-lbs分页标识，请在下一次翻页时作为入参传入
	PageResultKey string `json:"page_result_key,omitempty" xml:"page_result_key,omitempty"`
	// 搜索到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgMaterialOptionalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.PageResultKey = ""
	m.TotalResults = 0
}

var poolTaobaoTbkDgMaterialOptionalAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgMaterialOptionalAPIResponse)
	},
}

// GetTaobaoTbkDgMaterialOptionalAPIResponse 从 sync.Pool 获取 TaobaoTbkDgMaterialOptionalAPIResponse
func GetTaobaoTbkDgMaterialOptionalAPIResponse() *TaobaoTbkDgMaterialOptionalAPIResponse {
	return poolTaobaoTbkDgMaterialOptionalAPIResponse.Get().(*TaobaoTbkDgMaterialOptionalAPIResponse)
}

// ReleaseTaobaoTbkDgMaterialOptionalAPIResponse 将 TaobaoTbkDgMaterialOptionalAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgMaterialOptionalAPIResponse(v *TaobaoTbkDgMaterialOptionalAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgMaterialOptionalAPIResponse.Put(v)
}
