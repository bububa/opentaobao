package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tanx竞价失败反馈api APIRequest
taobao.tanx.biddingrefuses.get

竞价失败反馈根据创意id查询API提供
*/
type TaobaoTanxBiddingrefusesGetRequest struct {
    model.Params

    // dsp的创意id
    creativeIds   []string 

    // dsp在tanx的memberid
    memberId   int64 

    // dsp对应的tanx的token
    token   string 

    // 1970年到现在的毫秒
    signTime   int64 

    // 起始时间
    startTime   string 

    // 截止时间
    endTime   string 

}

func NewTaobaoTanxBiddingrefusesGetRequest() *TaobaoTanxBiddingrefusesGetRequest{
    return &TaobaoTanxBiddingrefusesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetApiMethodName() string {
    return "taobao.tanx.biddingrefuses.get"
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxBiddingrefusesGetRequest) SetCreativeIds(creativeIds []string) error {
    r.creativeIds = creativeIds
    r.Set("creative_ids", creativeIds)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetCreativeIds() []string {
    return r.creativeIds
}

func (r *TaobaoTanxBiddingrefusesGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxBiddingrefusesGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxBiddingrefusesGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxBiddingrefusesGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoTanxBiddingrefusesGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoTanxBiddingrefusesGetRequest) GetEndTime() string {
    return r.endTime
}

