package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivitytimeupdateAPIRequest 更新关联活动有效时间 API请求
// taobao.promotion.benefit.activity.time.update
//
// 更新关联权益的活动有效时间
type TaobaopromotionbenefitactivitytimeupdateAPIRequest struct {
	model.Params
	// 活动的开始时间
	_startTime string
	// 活动的i结束时间
	_endTime string
	// ISV活动关联权益后获得的关联ID
	_relationId int64
}

// NewTaobaopromotionbenefitactivitytimeupdateRequest 初始化TaobaopromotionbenefitactivitytimeupdateAPIRequest对象
func NewTaobaopromotionbenefitactivitytimeupdateRequest() *TaobaopromotionbenefitactivitytimeupdateAPIRequest {
	return &TaobaopromotionbenefitactivitytimeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.time.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 活动的开始时间
func (r *TaobaopromotionbenefitactivitytimeupdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动的i结束时间
func (r *TaobaopromotionbenefitactivitytimeupdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetRelationId is RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaopromotionbenefitactivitytimeupdateAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaopromotionbenefitactivitytimeupdateAPIRequest) GetRelationId() int64 {
	return r._relationId
}
