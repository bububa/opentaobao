package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tanx竞价失败反馈api API请求
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
type TaobaoTanxBiddingrefusesGetRequest struct {
    model.Params
    // dsp的创意id
    _creativeIds   []string
    // dsp在tanx的memberid
    _memberId   int64
    // dsp对应的tanx的token
    _token   string
    // 1970年到现在的毫秒
    _signTime   int64
    // 起始时间
    _startTime   string
    // 截止时间
    _endTime   string
}

// 初始化TaobaoTanxBiddingrefusesGetRequest对象
func NewTaobaoTanxBiddingrefusesGetRequest() *TaobaoTanxBiddingrefusesGetRequest{
    return &TaobaoTanxBiddingrefusesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxBiddingrefusesGetRequest) GetApiMethodName() string {
    return "taobao.tanx.biddingrefuses.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxBiddingrefusesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeIds Setter
// dsp的创意id
func (r *TaobaoTanxBiddingrefusesGetRequest) SetCreativeIds(_creativeIds []string) error {
    r._creativeIds = _creativeIds
    r.Set("creative_ids", _creativeIds)
    return nil
}

// CreativeIds Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetCreativeIds() []string {
    return r._creativeIds
}
// MemberId Setter
// dsp在tanx的memberid
func (r *TaobaoTanxBiddingrefusesGetRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp对应的tanx的token
func (r *TaobaoTanxBiddingrefusesGetRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 1970年到现在的毫秒
func (r *TaobaoTanxBiddingrefusesGetRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetSignTime() int64 {
    return r._signTime
}
// StartTime Setter
// 起始时间
func (r *TaobaoTanxBiddingrefusesGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 截止时间
func (r *TaobaoTanxBiddingrefusesGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoTanxBiddingrefusesGetRequest) GetEndTime() string {
    return r._endTime
}
