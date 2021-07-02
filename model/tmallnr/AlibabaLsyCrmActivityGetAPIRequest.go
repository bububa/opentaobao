package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityGetAPIRequest 私域导购查询活动详情 API请求
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
type AlibabaLsyCrmActivityGetAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 导购员id
	_guiderId int64
	// 摊位id
	_storeId int64
}

// NewAlibabaLsyCrmActivityGetRequest 初始化AlibabaLsyCrmActivityGetAPIRequest对象
func NewAlibabaLsyCrmActivityGetRequest() *AlibabaLsyCrmActivityGetAPIRequest {
	return &AlibabaLsyCrmActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is GuiderId Setter
// 导购员id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetGuiderId(_guiderId int64) error {
	r._guiderId = _guiderId
	r.Set("guider_id", _guiderId)
	return nil
}

// Get GuiderId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetGuiderId() int64 {
	return r._guiderId
}

// Set is StoreId Setter
// 摊位id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}
