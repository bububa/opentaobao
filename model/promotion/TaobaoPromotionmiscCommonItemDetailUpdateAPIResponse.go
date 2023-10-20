package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse 修改通用单品优惠详情 API返回值
// taobao.promotionmisc.common.item.detail.update
//
// 修改通用单品优惠详情。
// 1、该接口只修改活动下参与的商品的优惠信息，如需要增加、删除活动，请调用taobao.promotionmisc.common.item.activity.add和taobao.promotionmisc.common.item.activity.delete接口；
// 2、使用该接口时需要把未做修改的字段值也传入；
// 3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能修改
type TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemDetailUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemDetailUpdateAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemDetailUpdateAPIResponseModel is 修改通用单品优惠详情 成功返回结果
type TaobaoPromotionmiscCommonItemDetailUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_detail_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse
func GetTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse() *TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse {
	return poolTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse.Get().(*TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse 将 TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse(v *TaobaoPromotionmiscCommonItemDetailUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailUpdateAPIResponse.Put(v)
}
