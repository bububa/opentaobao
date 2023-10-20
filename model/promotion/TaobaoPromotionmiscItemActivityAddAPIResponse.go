package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityAddAPIResponse 创建无条件单品优惠活动 API返回值
// taobao.promotionmisc.item.activity.add
//
// 创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。&lt;br/&gt;2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。&lt;br/&gt;3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaoPromotionmiscItemActivityAddAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscItemActivityAddAPIResponseModel).Reset()
}

// TaobaoPromotionmiscItemActivityAddAPIResponseModel is 创建无条件单品优惠活动 成功返回结果
type TaobaoPromotionmiscItemActivityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动id。
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否保存成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ActivityId = 0
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscItemActivityAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscItemActivityAddAPIResponse)
	},
}

// GetTaobaoPromotionmiscItemActivityAddAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityAddAPIResponse
func GetTaobaoPromotionmiscItemActivityAddAPIResponse() *TaobaoPromotionmiscItemActivityAddAPIResponse {
	return poolTaobaoPromotionmiscItemActivityAddAPIResponse.Get().(*TaobaoPromotionmiscItemActivityAddAPIResponse)
}

// ReleaseTaobaoPromotionmiscItemActivityAddAPIResponse 将 TaobaoPromotionmiscItemActivityAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityAddAPIResponse(v *TaobaoPromotionmiscItemActivityAddAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityAddAPIResponse.Put(v)
}
