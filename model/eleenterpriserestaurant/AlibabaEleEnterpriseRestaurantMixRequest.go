package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
混合搜索店铺 APIRequest
alibaba.ele.enterprise.restaurant.mix

混合搜索店铺
*/
type AlibabaEleEnterpriseRestaurantMixRequest struct {
    model.Params

    // longitude和latitude用英文逗号分隔
    geo   string 

    // 首次查询无需传入，后续需要传入前次返回
    rankId   string 

    // 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
    start   int64 

    // 查询数量，默认是10，最大50
    limit   int64 

    // 人均消费金额上限，需要高于costFrom，不传表示不限
    costTo   int64 

    // 人均消费金额下限，最低为0，不传表示不限
    costFrom   int64 

    // 是否支持食安保（0-不限，1-支持食安保）不传表示不限
    insurance   int64 

    // 是否可开发票（0-不限，1-可开发票）不传表示不限
    invoice   int64 

    // 是否品牌商家（0-不限，1-品牌商家）不传表示不限
    isPremium   int64 

    // 是否新店（0-不限，1-新店）不传表示不限
    newRestaurant   int64 

    // 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
    deliveryMode   int64 

    // 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
    orderBy   int64 

    // 餐厅分类ids
    categoryIds   []int64 

    // 搜索关键词（关键字需要urlencode处理
    keyword   string 

    // 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
    isBookable   int64 

}

func NewAlibabaEleEnterpriseRestaurantMixRequest() *AlibabaEleEnterpriseRestaurantMixRequest{
    return &AlibabaEleEnterpriseRestaurantMixRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.mix"
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetGeo(geo string) error {
    r.geo = geo
    r.Set("geo", geo)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetGeo() string {
    return r.geo
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetRankId(rankId string) error {
    r.rankId = rankId
    r.Set("rank_id", rankId)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetRankId() string {
    return r.rankId
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetStart() int64 {
    return r.start
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetLimit() int64 {
    return r.limit
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetCostTo(costTo int64) error {
    r.costTo = costTo
    r.Set("cost_to", costTo)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetCostTo() int64 {
    return r.costTo
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetCostFrom(costFrom int64) error {
    r.costFrom = costFrom
    r.Set("cost_from", costFrom)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetCostFrom() int64 {
    return r.costFrom
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetInsurance(insurance int64) error {
    r.insurance = insurance
    r.Set("insurance", insurance)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetInsurance() int64 {
    return r.insurance
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetInvoice(invoice int64) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetInvoice() int64 {
    return r.invoice
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetIsPremium(isPremium int64) error {
    r.isPremium = isPremium
    r.Set("is_premium", isPremium)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetIsPremium() int64 {
    return r.isPremium
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetNewRestaurant(newRestaurant int64) error {
    r.newRestaurant = newRestaurant
    r.Set("new_restaurant", newRestaurant)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetNewRestaurant() int64 {
    return r.newRestaurant
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetDeliveryMode(deliveryMode int64) error {
    r.deliveryMode = deliveryMode
    r.Set("delivery_mode", deliveryMode)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetDeliveryMode() int64 {
    return r.deliveryMode
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetOrderBy(orderBy int64) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetOrderBy() int64 {
    return r.orderBy
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetCategoryIds(categoryIds []int64) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetCategoryIds() []int64 {
    return r.categoryIds
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlibabaEleEnterpriseRestaurantMixRequest) SetIsBookable(isBookable int64) error {
    r.isBookable = isBookable
    r.Set("is_bookable", isBookable)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMixRequest) GetIsBookable() int64 {
    return r.isBookable
}

