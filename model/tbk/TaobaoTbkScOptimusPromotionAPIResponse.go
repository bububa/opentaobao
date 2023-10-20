package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOptimusPromotionAPIResponse 淘宝客-服务商-权益物料精选 API返回值
// taobao.tbk.sc.optimus.promotion
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaoTbkScOptimusPromotionAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScOptimusPromotionAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScOptimusPromotionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScOptimusPromotionAPIResponseModel).Reset()
}

// TaobaoTbkScOptimusPromotionAPIResponseModel is 淘宝客-服务商-权益物料精选 成功返回结果
type TaobaoTbkScOptimusPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_optimus_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TaobaoTbkScOptimusPromotionMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_sc_optimus_promotion_map_data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScOptimusPromotionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTaobaoTbkScOptimusPromotionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScOptimusPromotionAPIResponse)
	},
}

// GetTaobaoTbkScOptimusPromotionAPIResponse 从 sync.Pool 获取 TaobaoTbkScOptimusPromotionAPIResponse
func GetTaobaoTbkScOptimusPromotionAPIResponse() *TaobaoTbkScOptimusPromotionAPIResponse {
	return poolTaobaoTbkScOptimusPromotionAPIResponse.Get().(*TaobaoTbkScOptimusPromotionAPIResponse)
}

// ReleaseTaobaoTbkScOptimusPromotionAPIResponse 将 TaobaoTbkScOptimusPromotionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScOptimusPromotionAPIResponse(v *TaobaoTbkScOptimusPromotionAPIResponse) {
	v.Reset()
	poolTaobaoTbkScOptimusPromotionAPIResponse.Put(v)
}
