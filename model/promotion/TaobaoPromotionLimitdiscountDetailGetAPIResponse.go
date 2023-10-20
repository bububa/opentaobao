package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionLimitdiscountDetailGetAPIResponse 限时打折详情查询 API返回值
// taobao.promotion.limitdiscount.detail.get
//
// 限时打折详情查询。查询出指定限时打折的对应商品记录信息。
type TaobaoPromotionLimitdiscountDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionLimitdiscountDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionLimitdiscountDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionLimitdiscountDetailGetAPIResponseModel).Reset()
}

// TaobaoPromotionLimitdiscountDetailGetAPIResponseModel is 限时打折详情查询 成功返回结果
type TaobaoPromotionLimitdiscountDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_limitdiscount_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 限时打折对应的商品详情列表。
	ItemDiscountDetailList []LimitDiscountDetail `json:"item_discount_detail_list,omitempty" xml:"item_discount_detail_list>limit_discount_detail,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionLimitdiscountDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemDiscountDetailList = m.ItemDiscountDetailList[:0]
}

var poolTaobaoPromotionLimitdiscountDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionLimitdiscountDetailGetAPIResponse)
	},
}

// GetTaobaoPromotionLimitdiscountDetailGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionLimitdiscountDetailGetAPIResponse
func GetTaobaoPromotionLimitdiscountDetailGetAPIResponse() *TaobaoPromotionLimitdiscountDetailGetAPIResponse {
	return poolTaobaoPromotionLimitdiscountDetailGetAPIResponse.Get().(*TaobaoPromotionLimitdiscountDetailGetAPIResponse)
}

// ReleaseTaobaoPromotionLimitdiscountDetailGetAPIResponse 将 TaobaoPromotionLimitdiscountDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionLimitdiscountDetailGetAPIResponse(v *TaobaoPromotionLimitdiscountDetailGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionLimitdiscountDetailGetAPIResponse.Put(v)
}
