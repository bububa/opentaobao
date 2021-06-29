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
    memberId   int64
    // dsp用户身份认证的TOKEN
    token   string
    // 当前时间戳，1970-01-01后的秒数
    signTime   int64
    // 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
    status   string
    // 分页的页码(第一页为1)
    page   int64
    // 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
    type   int64
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
func (r *TaobaoTanxCreativesGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxCreativesGetRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativesGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxCreativesGetRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativesGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxCreativesGetRequest) GetSignTime() int64 {
    return r.signTime
}
// Status Setter
// 创意的状态（全部ALL,通过PASS,拒绝REFUSE,未审核WAITING）
func (r *TaobaoTanxCreativesGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoTanxCreativesGetRequest) GetStatus() string {
    return r.status
}
// Page Setter
// 分页的页码(第一页为1)
func (r *TaobaoTanxCreativesGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r TaobaoTanxCreativesGetRequest) GetPage() int64 {
    return r.page
}
// Type Setter
// 所选创意的类型。1-->普通类型, 2-->视频贴片, 0 -->优先查询普通类型,无结果则查询视频贴片类型
func (r *TaobaoTanxCreativesGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoTanxCreativesGetRequest) GetType() int64 {
    return r.type
}
