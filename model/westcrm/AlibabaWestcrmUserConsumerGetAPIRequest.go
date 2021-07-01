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
type AlibabaWestcrmUserConsumerGetAPIRequest struct {
    model.Params
    // 园区id
    _campusId   int64
    // 用户id
    _ibUserId   int64
    // 开始时间
    _timeBegin   string
    // 结束时间
    _timeEnd   string
}

// 初始化AlibabaWestcrmUserConsumerGetAPIRequest对象
func NewAlibabaWestcrmUserConsumerGetRequest() *AlibabaWestcrmUserConsumerGetAPIRequest{
    return &AlibabaWestcrmUserConsumerGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.consumer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// IbUserId Setter
// 用户id
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetIbUserId(_ibUserId int64) error {
    r._ibUserId = _ibUserId
    r.Set("ib_user_id", _ibUserId)
    return nil
}

// IbUserId Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetIbUserId() int64 {
    return r._ibUserId
}
// TimeBegin Setter
// 开始时间
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetTimeBegin(_timeBegin string) error {
    r._timeBegin = _timeBegin
    r.Set("time_begin", _timeBegin)
    return nil
}

// TimeBegin Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetTimeBegin() string {
    return r._timeBegin
}
// TimeEnd Setter
// 结束时间
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetTimeEnd(_timeEnd string) error {
    r._timeEnd = _timeEnd
    r.Set("time_end", _timeEnd)
    return nil
}

// TimeEnd Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetTimeEnd() string {
    return r._timeEnd
}
