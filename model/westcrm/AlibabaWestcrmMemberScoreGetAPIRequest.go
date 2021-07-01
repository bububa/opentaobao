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
type AlibabaWestcrmMemberScoreGetAPIRequest struct {
    model.Params
    // requestId
    _requestId   string
    // 支付宝id
    _alipayId   string
    // 开始时间
    _startTime   string
    // 结束时间
    _endTime   string
}

// 初始化AlibabaWestcrmMemberScoreGetAPIRequest对象
func NewAlibabaWestcrmMemberScoreGetRequest() *AlibabaWestcrmMemberScoreGetAPIRequest{
    return &AlibabaWestcrmMemberScoreGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetApiMethodName() string {
    return "alibaba.westcrm.member.score.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestId Setter
// requestId
func (r *AlibabaWestcrmMemberScoreGetAPIRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetRequestId() string {
    return r._requestId
}
// AlipayId Setter
// 支付宝id
func (r *AlibabaWestcrmMemberScoreGetAPIRequest) SetAlipayId(_alipayId string) error {
    r._alipayId = _alipayId
    r.Set("alipay_id", _alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetAlipayId() string {
    return r._alipayId
}
// StartTime Setter
// 开始时间
func (r *AlibabaWestcrmMemberScoreGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间
func (r *AlibabaWestcrmMemberScoreGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaWestcrmMemberScoreGetAPIRequest) GetEndTime() string {
    return r._endTime
}
