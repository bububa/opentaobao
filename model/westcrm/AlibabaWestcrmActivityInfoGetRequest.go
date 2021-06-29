package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动详情 APIRequest
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

func NewAlibabaWestcrmActivityInfoGetRequest() *AlibabaWestcrmActivityInfoGetRequest{
    return &AlibabaWestcrmActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmActivityInfoGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.activity.info.get"
}

func (r AlibabaWestcrmActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmActivityInfoGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmActivityInfoGetRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmActivityInfoGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaWestcrmActivityInfoGetRequest) GetActivityId() int64 {
    return r.activityId
}

