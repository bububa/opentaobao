package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动详情 API请求
alibaba.westcrm.activity.info.get

根据id查询活动详情
*/
type AlibabaWestcrmActivityInfoGetRequest struct {
    model.Params
    // 园区id
    campusId   int64
    // 活动id
    activityId   int64
}

// 初始化AlibabaWestcrmActivityInfoGetRequest对象
func NewAlibabaWestcrmActivityInfoGetRequest() *AlibabaWestcrmActivityInfoGetRequest{
    return &AlibabaWestcrmActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmActivityInfoGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.activity.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmActivityInfoGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmActivityInfoGetRequest) GetCampusId() int64 {
    return r.campusId
}
// ActivityId Setter
// 活动id
func (r *AlibabaWestcrmActivityInfoGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaWestcrmActivityInfoGetRequest) GetActivityId() int64 {
    return r.activityId
}
