package eleenterpriserestaurant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantSearchAPIRequest 餐厅列表 API请求
// alibaba.ele.enterprise.restaurant.search
//
// 餐厅列表
type AlibabaEleEnterpriseRestaurantSearchAPIRequest struct {
	model.Params
	// 餐厅分类ids
	_categoryIds []int64
	// longitude和latitude用英文逗号分隔
	_geo string
	// 首次查询无需传入，后续需要传入前次返回
	_rankId string
	// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
	_isBookable string
	// 是否支持跨天预定 1:需要（不传该字段则不筛选）
	_crossDayBooking string
	// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
	_start int64
	// 查询数量，默认是10，最大50
	_limit int64
	// 人均消费金额上限，需要高于costFrom，不传表示不限
	_costTo int64
	// 人均消费金额下限，最低为0，不传表示不限
	_costFrom int64
	// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
	_insurance int64
	// 是否可开发票（0-不限，1-可开发票）不传表示不限
	_invoice int64
	// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
	_isPremium int64
	// 是否新店（0-不限，1-新店）不传表示不限
	_newRestaurant int64
	// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
	_deliveryMode int64
	// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
	_orderBy int64
}

// NewAlibabaEleEnterpriseRestaurantSearchRequest 初始化AlibabaEleEnterpriseRestaurantSearchAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantSearchRequest() *AlibabaEleEnterpriseRestaurantSearchAPIRequest {
	return &AlibabaEleEnterpriseRestaurantSearchAPIRequest{
		Params: model.NewParams(15),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) Reset() {
	r._categoryIds = r._categoryIds[:0]
	r._geo = ""
	r._rankId = ""
	r._isBookable = ""
	r._crossDayBooking = ""
	r._start = 0
	r._limit = 0
	r._costTo = 0
	r._costFrom = 0
	r._insurance = 0
	r._invoice = 0
	r._isPremium = 0
	r._newRestaurant = 0
	r._deliveryMode = 0
	r._orderBy = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryIds is CategoryIds Setter
// 餐厅分类ids
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetCategoryIds(_categoryIds []int64) error {
	r._categoryIds = _categoryIds
	r.Set("category_ids", _categoryIds)
	return nil
}

// GetCategoryIds CategoryIds Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetCategoryIds() []int64 {
	return r._categoryIds
}

// SetGeo is Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetGeo(_geo string) error {
	r._geo = _geo
	r.Set("geo", _geo)
	return nil
}

// GetGeo Geo Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetGeo() string {
	return r._geo
}

// SetRankId is RankId Setter
// 首次查询无需传入，后续需要传入前次返回
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetRankId(_rankId string) error {
	r._rankId = _rankId
	r.Set("rank_id", _rankId)
	return nil
}

// GetRankId RankId Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetRankId() string {
	return r._rankId
}

// SetIsBookable is IsBookable Setter
// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetIsBookable(_isBookable string) error {
	r._isBookable = _isBookable
	r.Set("is_bookable", _isBookable)
	return nil
}

// GetIsBookable IsBookable Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetIsBookable() string {
	return r._isBookable
}

// SetCrossDayBooking is CrossDayBooking Setter
// 是否支持跨天预定 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetCrossDayBooking(_crossDayBooking string) error {
	r._crossDayBooking = _crossDayBooking
	r.Set("cross_day_booking", _crossDayBooking)
	return nil
}

// GetCrossDayBooking CrossDayBooking Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetCrossDayBooking() string {
	return r._crossDayBooking
}

// SetStart is Start Setter
// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetLimit is Limit Setter
// 查询数量，默认是10，最大50
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetLimit() int64 {
	return r._limit
}

// SetCostTo is CostTo Setter
// 人均消费金额上限，需要高于costFrom，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetCostTo(_costTo int64) error {
	r._costTo = _costTo
	r.Set("cost_to", _costTo)
	return nil
}

// GetCostTo CostTo Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetCostTo() int64 {
	return r._costTo
}

// SetCostFrom is CostFrom Setter
// 人均消费金额下限，最低为0，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetCostFrom(_costFrom int64) error {
	r._costFrom = _costFrom
	r.Set("cost_from", _costFrom)
	return nil
}

// GetCostFrom CostFrom Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetCostFrom() int64 {
	return r._costFrom
}

// SetInsurance is Insurance Setter
// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetInsurance(_insurance int64) error {
	r._insurance = _insurance
	r.Set("insurance", _insurance)
	return nil
}

// GetInsurance Insurance Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetInsurance() int64 {
	return r._insurance
}

// SetInvoice is Invoice Setter
// 是否可开发票（0-不限，1-可开发票）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetInvoice(_invoice int64) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetInvoice() int64 {
	return r._invoice
}

// SetIsPremium is IsPremium Setter
// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetIsPremium(_isPremium int64) error {
	r._isPremium = _isPremium
	r.Set("is_premium", _isPremium)
	return nil
}

// GetIsPremium IsPremium Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetIsPremium() int64 {
	return r._isPremium
}

// SetNewRestaurant is NewRestaurant Setter
// 是否新店（0-不限，1-新店）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetNewRestaurant(_newRestaurant int64) error {
	r._newRestaurant = _newRestaurant
	r.Set("new_restaurant", _newRestaurant)
	return nil
}

// GetNewRestaurant NewRestaurant Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetNewRestaurant() int64 {
	return r._newRestaurant
}

// SetDeliveryMode is DeliveryMode Setter
// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetDeliveryMode(_deliveryMode int64) error {
	r._deliveryMode = _deliveryMode
	r.Set("delivery_mode", _deliveryMode)
	return nil
}

// GetDeliveryMode DeliveryMode Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetDeliveryMode() int64 {
	return r._deliveryMode
}

// SetOrderBy is OrderBy Setter
// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
func (r *AlibabaEleEnterpriseRestaurantSearchAPIRequest) SetOrderBy(_orderBy int64) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r AlibabaEleEnterpriseRestaurantSearchAPIRequest) GetOrderBy() int64 {
	return r._orderBy
}

var poolAlibabaEleEnterpriseRestaurantSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseRestaurantSearchRequest()
	},
}

// GetAlibabaEleEnterpriseRestaurantSearchRequest 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantSearchAPIRequest
func GetAlibabaEleEnterpriseRestaurantSearchAPIRequest() *AlibabaEleEnterpriseRestaurantSearchAPIRequest {
	return poolAlibabaEleEnterpriseRestaurantSearchAPIRequest.Get().(*AlibabaEleEnterpriseRestaurantSearchAPIRequest)
}

// ReleaseAlibabaEleEnterpriseRestaurantSearchAPIRequest 将 AlibabaEleEnterpriseRestaurantSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantSearchAPIRequest(v *AlibabaEleEnterpriseRestaurantSearchAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantSearchAPIRequest.Put(v)
}
