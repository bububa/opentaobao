package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户的消费总额 APIRequest
alibaba.westcrm.user.consumer.get

获取指定用户的消费总额
*/
type AlibabaWestcrmUserConsumerGetRequest struct {
    model.Params

    // 园区id
    campusId   int64 

    // 用户id
    ibUserId   int64 

    // 开始时间
    timeBegin   string 

    // 结束时间
    timeEnd   string 

}

func NewAlibabaWestcrmUserConsumerGetRequest() *AlibabaWestcrmUserConsumerGetRequest{
    return &AlibabaWestcrmUserConsumerGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.consumer.get"
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmUserConsumerGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmUserConsumerGetRequest) SetIbUserId(ibUserId int64) error {
    r.ibUserId = ibUserId
    r.Set("ib_user_id", ibUserId)
    return nil
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetIbUserId() int64 {
    return r.ibUserId
}

func (r *AlibabaWestcrmUserConsumerGetRequest) SetTimeBegin(timeBegin string) error {
    r.timeBegin = timeBegin
    r.Set("time_begin", timeBegin)
    return nil
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetTimeBegin() string {
    return r.timeBegin
}

func (r *AlibabaWestcrmUserConsumerGetRequest) SetTimeEnd(timeEnd string) error {
    r.timeEnd = timeEnd
    r.Set("time_end", timeEnd)
    return nil
}

func (r AlibabaWestcrmUserConsumerGetRequest) GetTimeEnd() string {
    return r.timeEnd
}

