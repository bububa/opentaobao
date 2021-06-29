package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
餐厅列表 API请求
alibaba.ele.enterprise.restaurant.search

餐厅列表
*/
type AlibabaEleEnterpriseRestaurantSearchRequest struct {
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
    // 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
    isBookable   string
    // 是否支持跨天预定 1:需要（不传该字段则不筛选）
    crossDayBooking   string
}

// 初始化AlibabaEleEnterpriseRestaurantSearchRequest对象
func NewAlibabaEleEnterpriseRestaurantSearchRequest() *AlibabaEleEnterpriseRestaurantSearchRequest{
    return &AlibabaEleEnterpriseRestaurantSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetGeo(geo string) error {
    r.geo = geo
    r.Set("geo", geo)
    return nil
}

// Geo Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetGeo() string {
    return r.geo
}
// RankId Setter
// 首次查询无需传入，后续需要传入前次返回
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetRankId(rankId string) error {
    r.rankId = rankId
    r.Set("rank_id", rankId)
    return nil
}

// RankId Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetRankId() string {
    return r.rankId
}
// Start Setter
// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetStart() int64 {
    return r.start
}
// Limit Setter
// 查询数量，默认是10，最大50
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetLimit() int64 {
    return r.limit
}
// CostTo Setter
// 人均消费金额上限，需要高于costFrom，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCostTo(costTo int64) error {
    r.costTo = costTo
    r.Set("cost_to", costTo)
    return nil
}

// CostTo Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCostTo() int64 {
    return r.costTo
}
// CostFrom Setter
// 人均消费金额下限，最低为0，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCostFrom(costFrom int64) error {
    r.costFrom = costFrom
    r.Set("cost_from", costFrom)
    return nil
}

// CostFrom Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCostFrom() int64 {
    return r.costFrom
}
// Insurance Setter
// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetInsurance(insurance int64) error {
    r.insurance = insurance
    r.Set("insurance", insurance)
    return nil
}

// Insurance Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetInsurance() int64 {
    return r.insurance
}
// Invoice Setter
// 是否可开发票（0-不限，1-可开发票）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetInvoice(invoice int64) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

// Invoice Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetInvoice() int64 {
    return r.invoice
}
// IsPremium Setter
// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetIsPremium(isPremium int64) error {
    r.isPremium = isPremium
    r.Set("is_premium", isPremium)
    return nil
}

// IsPremium Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetIsPremium() int64 {
    return r.isPremium
}
// NewRestaurant Setter
// 是否新店（0-不限，1-新店）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetNewRestaurant(newRestaurant int64) error {
    r.newRestaurant = newRestaurant
    r.Set("new_restaurant", newRestaurant)
    return nil
}

// NewRestaurant Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetNewRestaurant() int64 {
    return r.newRestaurant
}
// DeliveryMode Setter
// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetDeliveryMode(deliveryMode int64) error {
    r.deliveryMode = deliveryMode
    r.Set("delivery_mode", deliveryMode)
    return nil
}

// DeliveryMode Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetDeliveryMode() int64 {
    return r.deliveryMode
}
// OrderBy Setter
// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetOrderBy(orderBy int64) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetOrderBy() int64 {
    return r.orderBy
}
// CategoryIds Setter
// 餐厅分类ids
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCategoryIds(categoryIds []int64) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

// CategoryIds Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCategoryIds() []int64 {
    return r.categoryIds
}
// IsBookable Setter
// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetIsBookable(isBookable string) error {
    r.isBookable = isBookable
    r.Set("is_bookable", isBookable)
    return nil
}

// IsBookable Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetIsBookable() string {
    return r.isBookable
}
// CrossDayBooking Setter
// 是否支持跨天预定 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCrossDayBooking(crossDayBooking string) error {
    r.crossDayBooking = crossDayBooking
    r.Set("cross_day_booking", crossDayBooking)
    return nil
}

// CrossDayBooking Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCrossDayBooking() string {
    return r.crossDayBooking
}
