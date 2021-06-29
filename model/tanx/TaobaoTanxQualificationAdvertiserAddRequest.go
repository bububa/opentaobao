package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增广告主接口 API请求
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主
*/
type TaobaoTanxQualificationAdvertiserAddRequest struct {
    model.Params
    // 广告主对象
    _advertisers   []AdvertiserDto
    // dsp用户memberId
    _memberId   int64
    // dsp用户验证token
    _token   string
    // 从1970年到当前时间的秒
    _signTime   int64
}

// 初始化TaobaoTanxQualificationAdvertiserAddRequest对象
func NewTaobaoTanxQualificationAdvertiserAddRequest() *TaobaoTanxQualificationAdvertiserAddRequest{
    return &TaobaoTanxQualificationAdvertiserAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.advertiser.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Advertisers Setter
// 广告主对象
func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetAdvertisers(_advertisers []AdvertiserDto) error {
    r._advertisers = _advertisers
    r.Set("advertisers", _advertisers)
    return nil
}

// Advertisers Getter
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetAdvertisers() []AdvertiserDto {
    return r._advertisers
}
// MemberId Setter
// dsp用户memberId
func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户验证token
func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 从1970年到当前时间的秒
func (r *TaobaoTanxQualificationAdvertiserAddRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationAdvertiserAddRequest) GetSignTime() int64 {
    return r._signTime
}
