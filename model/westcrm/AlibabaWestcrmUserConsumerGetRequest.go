package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户的消费总额 API请求
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

// 初始化AlibabaWestcrmUserConsumerGetRequest对象
func NewAlibabaWestcrmUserConsumerGetRequest() *AlibabaWestcrmUserConsumerGetRequest{
    return &AlibabaWestcrmUserConsumerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserConsumerGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.consumer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserConsumerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmUserConsumerGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmUserConsumerGetRequest) GetCampusId() int64 {
    return r.campusId
}
// IbUserId Setter
// 用户id
func (r *AlibabaWestcrmUserConsumerGetRequest) SetIbUserId(ibUserId int64) error {
    r.ibUserId = ibUserId
    r.Set("ib_user_id", ibUserId)
    return nil
}

// IbUserId Getter
func (r AlibabaWestcrmUserConsumerGetRequest) GetIbUserId() int64 {
    return r.ibUserId
}
// TimeBegin Setter
// 开始时间
func (r *AlibabaWestcrmUserConsumerGetRequest) SetTimeBegin(timeBegin string) error {
    r.timeBegin = timeBegin
    r.Set("time_begin", timeBegin)
    return nil
}

// TimeBegin Getter
func (r AlibabaWestcrmUserConsumerGetRequest) GetTimeBegin() string {
    return r.timeBegin
}
// TimeEnd Setter
// 结束时间
func (r *AlibabaWestcrmUserConsumerGetRequest) SetTimeEnd(timeEnd string) error {
    r.timeEnd = timeEnd
    r.Set("time_end", timeEnd)
    return nil
}

// TimeEnd Getter
func (r AlibabaWestcrmUserConsumerGetRequest) GetTimeEnd() string {
    return r.timeEnd
}
