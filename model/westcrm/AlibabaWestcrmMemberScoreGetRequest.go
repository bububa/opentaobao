package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取会员某时间段积分 API请求
alibaba.westcrm.member.score.get

获取会员某时间段积分
*/
type AlibabaWestcrmMemberScoreGetRequest struct {
    model.Params
    // requestId
    requestId   string
    // 支付宝id
    alipayId   string
    // 开始时间
    startTime   string
    // 结束时间
    endTime   string
}

// 初始化AlibabaWestcrmMemberScoreGetRequest对象
func NewAlibabaWestcrmMemberScoreGetRequest() *AlibabaWestcrmMemberScoreGetRequest{
    return &AlibabaWestcrmMemberScoreGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMemberScoreGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.member.score.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMemberScoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestId Setter
// requestId
func (r *AlibabaWestcrmMemberScoreGetRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r AlibabaWestcrmMemberScoreGetRequest) GetRequestId() string {
    return r.requestId
}
// AlipayId Setter
// 支付宝id
func (r *AlibabaWestcrmMemberScoreGetRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaWestcrmMemberScoreGetRequest) GetAlipayId() string {
    return r.alipayId
}
// StartTime Setter
// 开始时间
func (r *AlibabaWestcrmMemberScoreGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaWestcrmMemberScoreGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *AlibabaWestcrmMemberScoreGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWestcrmMemberScoreGetRequest) GetEndTime() string {
    return r.endTime
}
