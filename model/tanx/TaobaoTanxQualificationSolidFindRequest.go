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
    _advertiserId   int64
    // 资质元素id列表
    _elementIds   []int64
    // dsp用户id
    _memberId   int64
    // dsp客户验证token
    _token   string
    // 1970年到现在的秒
    _signTime   int64
    // 查询起始页
    _page   int64
    // 分页大小
    _pageSize   int64
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
func (r *TaobaoTanxQualificationSolidFindRequest) SetAdvertiserId(_advertiserId int64) error {
    r._advertiserId = _advertiserId
    r.Set("advertiser_id", _advertiserId)
    return nil
}

// AdvertiserId Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetAdvertiserId() int64 {
    return r._advertiserId
}
// ElementIds Setter
// 资质元素id列表
func (r *TaobaoTanxQualificationSolidFindRequest) SetElementIds(_elementIds []int64) error {
    r._elementIds = _elementIds
    r.Set("element_ids", _elementIds)
    return nil
}

// ElementIds Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetElementIds() []int64 {
    return r._elementIds
}
// MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationSolidFindRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp客户验证token
func (r *TaobaoTanxQualificationSolidFindRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 1970年到现在的秒
func (r *TaobaoTanxQualificationSolidFindRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetSignTime() int64 {
    return r._signTime
}
// Page Setter
// 查询起始页
func (r *TaobaoTanxQualificationSolidFindRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 分页大小
func (r *TaobaoTanxQualificationSolidFindRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTanxQualificationSolidFindRequest) GetPageSize() int64 {
    return r._pageSize
}
