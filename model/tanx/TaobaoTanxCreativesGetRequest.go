package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取DSP用户的创意审核结果 API请求
taobao.tanx.creatives.get

批量获取DSP用户的创意审核结果
*/
type TaobaoTanxCreativesGetRequest struct {
    model.Params
    // DSP的memberId
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
    // 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
    _status   string
    // 分页的页码(第一页为1)
    _page   int64
    // 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
    _type   int64
}

// 初始化TaobaoTanxCreativesGetRequest对象
func NewTaobaoTanxCreativesGetRequest() *TaobaoTanxCreativesGetRequest{
    return &TaobaoTanxCreativesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativesGetRequest) GetApiMethodName() string {
    return "taobao.tanx.creatives.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativesGetRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxCreativesGetRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativesGetRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxCreativesGetRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativesGetRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxCreativesGetRequest) GetSignTime() int64 {
    return r._signTime
}
// Status Setter
// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
func (r *TaobaoTanxCreativesGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoTanxCreativesGetRequest) GetStatus() string {
    return r._status
}
// Page Setter
// 分页的页码(第一页为1)
func (r *TaobaoTanxCreativesGetRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoTanxCreativesGetRequest) GetPage() int64 {
    return r._page
}
// Type Setter
// 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
func (r *TaobaoTanxCreativesGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoTanxCreativesGetRequest) GetType() int64 {
    return r._type
}
