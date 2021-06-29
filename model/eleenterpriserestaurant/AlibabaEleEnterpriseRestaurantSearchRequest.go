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
    _geo   string
    // 首次查询无需传入，后续需要传入前次返回
    _rankId   string
    // 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
    _start   int64
    // 查询数量，默认是10，最大50
    _limit   int64
    // 人均消费金额上限，需要高于costFrom，不传表示不限
    _costTo   int64
    // 人均消费金额下限，最低为0，不传表示不限
    _costFrom   int64
    // 是否支持食安保（0-不限，1-支持食安保）不传表示不限
    _insurance   int64
    // 是否可开发票（0-不限，1-可开发票）不传表示不限
    _invoice   int64
    // 是否品牌商家（0-不限，1-品牌商家）不传表示不限
    _isPremium   int64
    // 是否新店（0-不限，1-新店）不传表示不限
    _newRestaurant   int64
    // 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
    _deliveryMode   int64
    // 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
    _orderBy   int64
    // 餐厅分类ids
    _categoryIds   []int64
    // 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
    _isBookable   string
    // 是否支持跨天预定 1:需要（不传该字段则不筛选）
    _crossDayBooking   string
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
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetGeo(_geo string) error {
    r._geo = _geo
    r.Set("geo", _geo)
    return nil
}

// Geo Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetGeo() string {
    return r._geo
}
// RankId Setter
// 首次查询无需传入，后续需要传入前次返回
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetRankId(_rankId string) error {
    r._rankId = _rankId
    r.Set("rank_id", _rankId)
    return nil
}

// RankId Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetRankId() string {
    return r._rankId
}
// Start Setter
// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetStart() int64 {
    return r._start
}
// Limit Setter
// 查询数量，默认是10，最大50
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetLimit() int64 {
    return r._limit
}
// CostTo Setter
// 人均消费金额上限，需要高于costFrom，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCostTo(_costTo int64) error {
    r._costTo = _costTo
    r.Set("cost_to", _costTo)
    return nil
}

// CostTo Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCostTo() int64 {
    return r._costTo
}
// CostFrom Setter
// 人均消费金额下限，最低为0，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCostFrom(_costFrom int64) error {
    r._costFrom = _costFrom
    r.Set("cost_from", _costFrom)
    return nil
}

// CostFrom Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCostFrom() int64 {
    return r._costFrom
}
// Insurance Setter
// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetInsurance(_insurance int64) error {
    r._insurance = _insurance
    r.Set("insurance", _insurance)
    return nil
}

// Insurance Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetInsurance() int64 {
    return r._insurance
}
// Invoice Setter
// 是否可开发票（0-不限，1-可开发票）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetInvoice(_invoice int64) error {
    r._invoice = _invoice
    r.Set("invoice", _invoice)
    return nil
}

// Invoice Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetInvoice() int64 {
    return r._invoice
}
// IsPremium Setter
// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetIsPremium(_isPremium int64) error {
    r._isPremium = _isPremium
    r.Set("is_premium", _isPremium)
    return nil
}

// IsPremium Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetIsPremium() int64 {
    return r._isPremium
}
// NewRestaurant Setter
// 是否新店（0-不限，1-新店）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetNewRestaurant(_newRestaurant int64) error {
    r._newRestaurant = _newRestaurant
    r.Set("new_restaurant", _newRestaurant)
    return nil
}

// NewRestaurant Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetNewRestaurant() int64 {
    return r._newRestaurant
}
// DeliveryMode Setter
// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetDeliveryMode(_deliveryMode int64) error {
    r._deliveryMode = _deliveryMode
    r.Set("delivery_mode", _deliveryMode)
    return nil
}

// DeliveryMode Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetDeliveryMode() int64 {
    return r._deliveryMode
}
// OrderBy Setter
// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetOrderBy(_orderBy int64) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetOrderBy() int64 {
    return r._orderBy
}
// CategoryIds Setter
// 餐厅分类ids
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCategoryIds(_categoryIds []int64) error {
    r._categoryIds = _categoryIds
    r.Set("category_ids", _categoryIds)
    return nil
}

// CategoryIds Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCategoryIds() []int64 {
    return r._categoryIds
}
// IsBookable Setter
// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetIsBookable(_isBookable string) error {
    r._isBookable = _isBookable
    r.Set("is_bookable", _isBookable)
    return nil
}

// IsBookable Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetIsBookable() string {
    return r._isBookable
}
// CrossDayBooking Setter
// 是否支持跨天预定 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchRequest) SetCrossDayBooking(_crossDayBooking string) error {
    r._crossDayBooking = _crossDayBooking
    r.Set("cross_day_booking", _crossDayBooking)
    return nil
}

// CrossDayBooking Getter
func (r AlibabaEleEnterpriseRestaurantSearchRequest) GetCrossDayBooking() string {
    return r._crossDayBooking
}
