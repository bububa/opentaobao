package tmallnr

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLsyCrmActivityGetAPIRequest) Reset() {
	r._activityId = 0
	r._guiderId = 0
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetGuiderId is GuiderId Setter
// 导购员id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetGuiderId(_guiderId int64) error {
	r._guiderId = _guiderId
	r.Set("guider_id", _guiderId)
	return nil
}

// GetGuiderId GuiderId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetGuiderId() int64 {
	return r._guiderId
}

// SetStoreId is StoreId Setter
// 摊位id
func (r *AlibabaLsyCrmActivityGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaLsyCrmActivityGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolAlibabaLsyCrmActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLsyCrmActivityGetRequest()
	},
}

// GetAlibabaLsyCrmActivityGetRequest 从 sync.Pool 获取 AlibabaLsyCrmActivityGetAPIRequest
func GetAlibabaLsyCrmActivityGetAPIRequest() *AlibabaLsyCrmActivityGetAPIRequest {
	return poolAlibabaLsyCrmActivityGetAPIRequest.Get().(*AlibabaLsyCrmActivityGetAPIRequest)
}

// ReleaseAlibabaLsyCrmActivityGetAPIRequest 将 AlibabaLsyCrmActivityGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaLsyCrmActivityGetAPIRequest(v *AlibabaLsyCrmActivityGetAPIRequest) {
	v.Reset()
	poolAlibabaLsyCrmActivityGetAPIRequest.Put(v)
}
