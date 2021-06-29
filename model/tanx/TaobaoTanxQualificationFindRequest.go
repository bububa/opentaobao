package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资质查询接口 API请求
taobao.tanx.qualification.find

资质查询接口
*/
type TaobaoTanxQualificationFindRequest struct {
    model.Params
    // dsp客户在tanx的memberId
    memberId   int64
    // dsp客户在tanx签名的token
    token   string
    // 当前时间,从1970年算的毫秒
    signTime   int64
    // 分页的第几页，从1开始
    page   int64
    // 分页大小，限制200
    pageSize   int64
    // 广告主资质查询dto
    query   *QualificationQuery
}

// 初始化TaobaoTanxQualificationFindRequest对象
func NewTaobaoTanxQualificationFindRequest() *TaobaoTanxQualificationFindRequest{
    return &TaobaoTanxQualificationFindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationFindRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// dsp客户在tanx的memberId
func (r *TaobaoTanxQualificationFindRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationFindRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp客户在tanx签名的token
func (r *TaobaoTanxQualificationFindRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationFindRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 当前时间,从1970年算的毫秒
func (r *TaobaoTanxQualificationFindRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationFindRequest) GetSignTime() int64 {
    return r.signTime
}
// Page Setter
// 分页的第几页，从1开始
func (r *TaobaoTanxQualificationFindRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r TaobaoTanxQualificationFindRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 分页大小，限制200
func (r *TaobaoTanxQualificationFindRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTanxQualificationFindRequest) GetPageSize() int64 {
    return r.pageSize
}
// Query Setter
// 广告主资质查询dto
func (r *TaobaoTanxQualificationFindRequest) SetQuery(query *QualificationQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TaobaoTanxQualificationFindRequest) GetQuery() *QualificationQuery {
    return r.query
}
