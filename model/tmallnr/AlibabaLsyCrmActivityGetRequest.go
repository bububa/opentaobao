package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购查询活动详情 API请求
alibaba.lsy.crm.activity.get

私域导购查询活动详情
*/
type AlibabaLsyCrmActivityGetRequest struct {
    model.Params
    // 活动id
    _activityId   int64
    // 导购员id
    _guiderId   int64
    // 摊位id
    _storeId   int64
}

// 初始化AlibabaLsyCrmActivityGetRequest对象
func NewAlibabaLsyCrmActivityGetRequest() *AlibabaLsyCrmActivityGetRequest{
    return &AlibabaLsyCrmActivityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id
func (r *AlibabaLsyCrmActivityGetRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetActivityId() int64 {
    return r._activityId
}
// GuiderId Setter
// 导购员id
func (r *AlibabaLsyCrmActivityGetRequest) SetGuiderId(_guiderId int64) error {
    r._guiderId = _guiderId
    r.Set("guider_id", _guiderId)
    return nil
}

// GuiderId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetGuiderId() int64 {
    return r._guiderId
}
// StoreId Setter
// 摊位id
func (r *AlibabaLsyCrmActivityGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetStoreId() int64 {
    return r._storeId
}
