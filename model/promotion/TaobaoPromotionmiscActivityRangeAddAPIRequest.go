package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeAddAPIRequest 增加活动参与的商品 API请求
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
type TaobaoPromotionmiscActivityRangeAddAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
	// 商品id,多个id用逗号隔开，一次不超过50个。
	_ids string
}

// NewTaobaoPromotionmiscActivityRangeAddRequest 初始化TaobaoPromotionmiscActivityRangeAddAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeAddRequest() *TaobaoPromotionmiscActivityRangeAddAPIRequest {
	return &TaobaoPromotionmiscActivityRangeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAddAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is Ids Setter
// 商品id,多个id用逗号隔开，一次不超过50个。
func (r *TaobaoPromotionmiscActivityRangeAddAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// Get Ids Getter
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetIds() string {
	return r._ids
}
