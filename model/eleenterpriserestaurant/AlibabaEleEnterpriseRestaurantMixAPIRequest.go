package eleenterpriserestaurant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantMixAPIRequest 混合搜索店铺 API请求
// alibaba.ele.enterprise.restaurant.mix
//
// 混合搜索店铺
type AlibabaEleEnterpriseRestaurantMixAPIRequest struct {
	model.Params
	// 餐厅分类ids
	_categoryIds []int64
	// longitude和latitude用英文逗号分隔
	_geo string
	// 首次查询无需传入，后续需要传入前次返回
	_rankId string
	// 搜索关键词（关键字需要urlencode处理
	_keyword string
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
	// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
	_isBookable int64
}

// NewAlibabaEleEnterpriseRestaurantMixRequest 初始化AlibabaEleEnterpriseRestaurantMixAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantMixRequest() *AlibabaEleEnterpriseRestaurantMixAPIRequest {
	return &AlibabaEleEnterpriseRestaurantMixAPIRequest{
		Params: model.NewParams(15),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) Reset() {
	r._categoryIds = r._categoryIds[:0]
	r._geo = ""
	r._rankId = ""
	r._keyword = ""
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
	r._isBookable = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.mix"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryIds is CategoryIds Setter
// 餐厅分类ids
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetCategoryIds(_categoryIds []int64) error {
	r._categoryIds = _categoryIds
	r.Set("category_ids", _categoryIds)
	return nil
}

// GetCategoryIds CategoryIds Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetCategoryIds() []int64 {
	return r._categoryIds
}

// SetGeo is Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetGeo(_geo string) error {
	r._geo = _geo
	r.Set("geo", _geo)
	return nil
}

// GetGeo Geo Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetGeo() string {
	return r._geo
}

// SetRankId is RankId Setter
// 首次查询无需传入，后续需要传入前次返回
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetRankId(_rankId string) error {
	r._rankId = _rankId
	r.Set("rank_id", _rankId)
	return nil
}

// GetRankId RankId Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetRankId() string {
	return r._rankId
}

// SetKeyword is Keyword Setter
// 搜索关键词（关键字需要urlencode处理
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetStart is Start Setter
// 查询起始位置，默认为0。如果传的是10，那么餐厅会从第11个开始返回
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetStart() int64 {
	return r._start
}

// SetLimit is Limit Setter
// 查询数量，默认是10，最大50
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetLimit() int64 {
	return r._limit
}

// SetCostTo is CostTo Setter
// 人均消费金额上限，需要高于costFrom，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetCostTo(_costTo int64) error {
	r._costTo = _costTo
	r.Set("cost_to", _costTo)
	return nil
}

// GetCostTo CostTo Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetCostTo() int64 {
	return r._costTo
}

// SetCostFrom is CostFrom Setter
// 人均消费金额下限，最低为0，不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetCostFrom(_costFrom int64) error {
	r._costFrom = _costFrom
	r.Set("cost_from", _costFrom)
	return nil
}

// GetCostFrom CostFrom Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetCostFrom() int64 {
	return r._costFrom
}

// SetInsurance is Insurance Setter
// 是否支持食安保（0-不限，1-支持食安保）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetInsurance(_insurance int64) error {
	r._insurance = _insurance
	r.Set("insurance", _insurance)
	return nil
}

// GetInsurance Insurance Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetInsurance() int64 {
	return r._insurance
}

// SetInvoice is Invoice Setter
// 是否可开发票（0-不限，1-可开发票）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetInvoice(_invoice int64) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetInvoice() int64 {
	return r._invoice
}

// SetIsPremium is IsPremium Setter
// 是否品牌商家（0-不限，1-品牌商家）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetIsPremium(_isPremium int64) error {
	r._isPremium = _isPremium
	r.Set("is_premium", _isPremium)
	return nil
}

// GetIsPremium IsPremium Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetIsPremium() int64 {
	return r._isPremium
}

// SetNewRestaurant is NewRestaurant Setter
// 是否新店（0-不限，1-新店）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetNewRestaurant(_newRestaurant int64) error {
	r._newRestaurant = _newRestaurant
	r.Set("new_restaurant", _newRestaurant)
	return nil
}

// GetNewRestaurant NewRestaurant Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetNewRestaurant() int64 {
	return r._newRestaurant
}

// SetDeliveryMode is DeliveryMode Setter
// 配送方式（0-不限， 1-蜂鸟专送）不传表示不限
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetDeliveryMode(_deliveryMode int64) error {
	r._deliveryMode = _deliveryMode
	r.Set("delivery_mode", _deliveryMode)
	return nil
}

// GetDeliveryMode DeliveryMode Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetDeliveryMode() int64 {
	return r._deliveryMode
}

// SetOrderBy is OrderBy Setter
// 排序选项（1-默认排序（热门）， 2-评价星级由高到低， 3-起送价由低到高， 4-销量由高到低， 5-送餐速度由快到慢， 6-餐厅距离由近到远， 7-订单量由高到低）
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetOrderBy(_orderBy int64) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetOrderBy() int64 {
	return r._orderBy
}

// SetIsBookable is IsBookable Setter
// 是否筛选支持预定 0:不需要 1:需要（不传该字段则不筛选）
func (r *AlibabaEleEnterpriseRestaurantMixAPIRequest) SetIsBookable(_isBookable int64) error {
	r._isBookable = _isBookable
	r.Set("is_bookable", _isBookable)
	return nil
}

// GetIsBookable IsBookable Getter
func (r AlibabaEleEnterpriseRestaurantMixAPIRequest) GetIsBookable() int64 {
	return r._isBookable
}

var poolAlibabaEleEnterpriseRestaurantMixAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseRestaurantMixRequest()
	},
}

// GetAlibabaEleEnterpriseRestaurantMixRequest 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantMixAPIRequest
func GetAlibabaEleEnterpriseRestaurantMixAPIRequest() *AlibabaEleEnterpriseRestaurantMixAPIRequest {
	return poolAlibabaEleEnterpriseRestaurantMixAPIRequest.Get().(*AlibabaEleEnterpriseRestaurantMixAPIRequest)
}

// ReleaseAlibabaEleEnterpriseRestaurantMixAPIRequest 将 AlibabaEleEnterpriseRestaurantMixAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantMixAPIRequest(v *AlibabaEleEnterpriseRestaurantMixAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantMixAPIRequest.Put(v)
}
