package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增广告主接口 APIRequest
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
*/
type TaobaoTanxQualificationAdvertiserAddRequest struct {
    model.Params

    // 广告主对象
    advertisers   []AdvertiserDto 

    // dsp用户memberId
    memberId   int64 

    // dsp用户验证token
    token   string 

    // 从1970年到当前时间的秒
    signTime   int64 

}

func NewTaobaoTanxQualificationAdvertiserAddRequest() *TaobaoTanxQualificationAdvertiserAddRequest{
    return &TaobaoTanxQualificationAdvertiserAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.advertiser.add"
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetAdvertisers(advertisers []AdvertiserDto) error {
    r.advertisers = advertisers
    r.Set("advertisers", advertisers)
    return nil
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetAdvertisers() []AdvertiserDto {
    return r.advertisers
}

func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxQualificationAdvertiserAddRequest) GetSignTime() int64 {
    return r.signTime
}

