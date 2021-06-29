package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户固态共享资质 API请求
taobao.tanx.qualification.solid.find

接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
*/
type TaobaoTanxQualificationSolidFindRequest struct {
    model.Params
    // 广告主id
    advertiserId   int64
    // 资质元素id列表
    elementIds   []int64
    // dsp用户id
    memberId   int64
    // dsp客户验证token
    token   string
    // 1970年到现在的秒
    signTime   int64
    // 查询起始页
    page   int64
    // 分页大小
    pageSize   int64
}

// 初始化TaobaoTanxQualificationSolidFindRequest对象
func NewTaobaoTanxQualificationSolidFindRequest() *TaobaoTanxQualificationSolidFindRequest{
    return &TaobaoTanxQualificationSolidFindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationSolidFindRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.solid.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationSolidFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdvertiserId Setter
// 广告主id
func (r *TaobaoTanxQualificationSolidFindRequest) SetAdvertiserId(advertiserId int64) error {
    r.advertiserId = advertiserId
    r.Set("advertiser_id", advertiserId)
    return nil
}

// AdvertiserId Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetAdvertiserId() int64 {
    return r.advertiserId
}
// ElementIds Setter
// 资质元素id列表
func (r *TaobaoTanxQualificationSolidFindRequest) SetElementIds(elementIds []int64) error {
    r.elementIds = elementIds
    r.Set("element_ids", elementIds)
    return nil
}

// ElementIds Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetElementIds() []int64 {
    return r.elementIds
}
// MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationSolidFindRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp客户验证token
func (r *TaobaoTanxQualificationSolidFindRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 1970年到现在的秒
func (r *TaobaoTanxQualificationSolidFindRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetSignTime() int64 {
    return r.signTime
}
// Page Setter
// 查询起始页
func (r *TaobaoTanxQualificationSolidFindRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 分页大小
func (r *TaobaoTanxQualificationSolidFindRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetPageSize() int64 {
    return r.pageSize
}
