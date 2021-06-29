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
    activityId   int64
    // 导购员id
    guiderId   int64
    // 摊位id
    storeId   int64
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
func (r *AlibabaLsyCrmActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}
// GuiderId Setter
// 导购员id
func (r *AlibabaLsyCrmActivityGetRequest) SetGuiderId(guiderId int64) error {
    r.guiderId = guiderId
    r.Set("guider_id", guiderId)
    return nil
}

// GuiderId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetGuiderId() int64 {
    return r.guiderId
}
// StoreId Setter
// 摊位id
func (r *AlibabaLsyCrmActivityGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaLsyCrmActivityGetRequest) GetStoreId() int64 {
    return r.storeId
}
