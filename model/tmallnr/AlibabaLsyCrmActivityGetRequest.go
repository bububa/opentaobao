package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购查询活动详情 APIRequest
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

func NewAlibabaLsyCrmActivityGetRequest() *AlibabaLsyCrmActivityGetRequest{
    return &AlibabaLsyCrmActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmActivityGetRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.get"
}

func (r AlibabaLsyCrmActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaLsyCrmActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *AlibabaLsyCrmActivityGetRequest) SetGuiderId(guiderId int64) error {
    r.guiderId = guiderId
    r.Set("guider_id", guiderId)
    return nil
}

func (r AlibabaLsyCrmActivityGetRequest) GetGuiderId() int64 {
    return r.guiderId
}

func (r *AlibabaLsyCrmActivityGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaLsyCrmActivityGetRequest) GetStoreId() int64 {
    return r.storeId
}

