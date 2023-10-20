package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityTimeUpdateAPIRequest 更新关联活动有效时间 API请求
// taobao.promotion.benefit.activity.time.update
//
// 更新关联权益的活动有效时间
type TaobaoPromotionBenefitActivityTimeUpdateAPIRequest struct {
	model.Params
	// 活动的开始时间
	_startTime string
	// 活动的i结束时间
	_endTime string
	// ISV活动关联权益后获得的关联ID
	_relationId int64
}

// NewTaobaoPromotionBenefitActivityTimeUpdateRequest 初始化TaobaoPromotionBenefitActivityTimeUpdateAPIRequest对象
func NewTaobaoPromotionBenefitActivityTimeUpdateRequest() *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest {
	return &TaobaoPromotionBenefitActivityTimeUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) Reset() {
	r._startTime = ""
	r._endTime = ""
	r._relationId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.time.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 活动的开始时间
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动的i结束时间
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetRelationId is RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetRelationId() int64 {
	return r._relationId
}

var poolTaobaoPromotionBenefitActivityTimeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionBenefitActivityTimeUpdateRequest()
	},
}

// GetTaobaoPromotionBenefitActivityTimeUpdateRequest 从 sync.Pool 获取 TaobaoPromotionBenefitActivityTimeUpdateAPIRequest
func GetTaobaoPromotionBenefitActivityTimeUpdateAPIRequest() *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest {
	return poolTaobaoPromotionBenefitActivityTimeUpdateAPIRequest.Get().(*TaobaoPromotionBenefitActivityTimeUpdateAPIRequest)
}

// ReleaseTaobaoPromotionBenefitActivityTimeUpdateAPIRequest 将 TaobaoPromotionBenefitActivityTimeUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityTimeUpdateAPIRequest(v *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityTimeUpdateAPIRequest.Put(v)
}
