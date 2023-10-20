package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitygetAPIRequest 私域导购查询活动详情 API请求
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
type AlibabalsycrmactivitygetAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 导购员id
	_guiderId int64
	// 摊位id
	_storeId int64
}

// NewAlibabalsycrmactivitygetRequest 初始化AlibabalsycrmactivitygetAPIRequest对象
func NewAlibabalsycrmactivitygetRequest() *AlibabalsycrmactivitygetAPIRequest {
	return &AlibabalsycrmactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitygetAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabalsycrmactivitygetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabalsycrmactivitygetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetGuiderId is GuiderId Setter
// 导购员id
func (r *AlibabalsycrmactivitygetAPIRequest) SetGuiderId(_guiderId int64) error {
	r._guiderId = _guiderId
	r.Set("guider_id", _guiderId)
	return nil
}

// GetGuiderId GuiderId Getter
func (r AlibabalsycrmactivitygetAPIRequest) GetGuiderId() int64 {
	return r._guiderId
}

// SetStoreId is StoreId Setter
// 摊位id
func (r *AlibabalsycrmactivitygetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabalsycrmactivitygetAPIRequest) GetStoreId() int64 {
	return r._storeId
}
