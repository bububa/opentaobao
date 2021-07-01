package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscMjsActivityGetAPIRequest
查询满就送活动 API请求
taobao.promotionmisc.mjs.activity.get

查询满就送活动 */
type TaobaoPromotionmiscMjsActivityGetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscMjsActivityGetRequest 初始化TaobaoPromotionmiscMjsActivityGetAPIRequest对象
func NewTaobaoPromotionmiscMjsActivityGetRequest() *TaobaoPromotionmiscMjsActivityGetAPIRequest {
	return &TaobaoPromotionmiscMjsActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
