package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户固态共享资质 APIRequest
taobao.tanx.qualification.solid.find

接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
*/
type TaobaoTanxQualificationSolidFindRequest struct {
    model.Params

    // 广告主id
    advertiserId   int64 

    // 资质元素id列表
    elementIds   []Number 

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

func NewTaobaoTanxQualificationSolidFindRequest() *TaobaoTanxQualificationSolidFindRequest{
    return &TaobaoTanxQualificationSolidFindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxQualificationSolidFindRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.solid.find"
}

func (r TaobaoTanxQualificationSolidFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxQualificationSolidFindRequest) SetAdvertiserId(advertiserId int64) error {
    r.advertiserId = advertiserId
    r.Set("advertiser_id", advertiserId)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetAdvertiserId() int64 {
    return r.advertiserId
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetElementIds(elementIds []Number) error {
    r.elementIds = elementIds
    r.Set("element_ids", elementIds)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetElementIds() []Number {
    return r.elementIds
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoTanxQualificationSolidFindRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTanxQualificationSolidFindRequest) GetPageSize() int64 {
    return r.pageSize
}

