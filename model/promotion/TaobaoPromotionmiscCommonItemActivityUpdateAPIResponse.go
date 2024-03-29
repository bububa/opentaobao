package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse 修改通用单品优惠活动 API返回值
// taobao.promotionmisc.common.item.activity.update
//
// 修改通用单品优惠活动。
// 1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
// 2、使用该接口时需要把未做修改的字段值也传入
type TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemActivityUpdateAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemActivityUpdateAPIResponseModel is 修改通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse
func GetTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse() *TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse {
	return poolTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse.Get().(*TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse 将 TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse(v *TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityUpdateAPIResponse.Put(v)
}
