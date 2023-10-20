package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgOptimusPromotionAPIResponse 淘宝客-推广者-权益物料精选 API返回值
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaoTbkDgOptimusPromotionAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgOptimusPromotionAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgOptimusPromotionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgOptimusPromotionAPIResponseModel).Reset()
}

// TaobaoTbkDgOptimusPromotionAPIResponseModel is 淘宝客-推广者-权益物料精选 成功返回结果
type TaobaoTbkDgOptimusPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_optimus_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaoTbkDgOptimusPromotionMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_dg_optimus_promotion_map_data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgOptimusPromotionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTaobaoTbkDgOptimusPromotionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgOptimusPromotionAPIResponse)
	},
}

// GetTaobaoTbkDgOptimusPromotionAPIResponse 从 sync.Pool 获取 TaobaoTbkDgOptimusPromotionAPIResponse
func GetTaobaoTbkDgOptimusPromotionAPIResponse() *TaobaoTbkDgOptimusPromotionAPIResponse {
	return poolTaobaoTbkDgOptimusPromotionAPIResponse.Get().(*TaobaoTbkDgOptimusPromotionAPIResponse)
}

// ReleaseTaobaoTbkDgOptimusPromotionAPIResponse 将 TaobaoTbkDgOptimusPromotionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgOptimusPromotionAPIResponse(v *TaobaoTbkDgOptimusPromotionAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgOptimusPromotionAPIResponse.Put(v)
}
