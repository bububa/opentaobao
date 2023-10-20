package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialitemsearchAPIRequest AE社交选品 API请求
// aliexpress.social.item.search
//
// AE社交选品,通过各种筛选条件对社交商品池进行筛选
type AliexpresssocialitemsearchAPIRequest struct {
	model.Params
	// order by properties
	_orderBy string
	// 佣金最大值
	_commissionRateMax string
	// 佣金最小值
	_commissionRateMin string
	// 最低价格
	_minPrice string
	// 最高价格
	_maxPrice string
	// 搜索关键字
	_keyword string
	// shipTo国家
	_shipTo string
	// 评价分数
	_commentScore string
	// 币种
	_currency string
	// locale,格式为language+"_"+country
	_locale string
	// page size
	_pageSize int64
	// 物流时效
	_logisticsTime int64
	// 类目ID
	_cateId int64
	// 页码
	_pageNo int64
	// 是否有视频
	_hasVideo bool
	// 是否逆序
	_desc bool
	// 是否免邮
	_isShipFree bool
	// 是否联盟商品
	_allianceItem bool
}

// NewAliexpresssocialitemsearchRequest 初始化AliexpresssocialitemsearchAPIRequest对象
func NewAliexpresssocialitemsearchRequest() *AliexpresssocialitemsearchAPIRequest {
	return &AliexpresssocialitemsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialitemsearchAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.item.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialitemsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialitemsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderBy is OrderBy Setter
// order by properties
func (r *AliexpresssocialitemsearchAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r AliexpresssocialitemsearchAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetCommissionRateMax is CommissionRateMax Setter
// 佣金最大值
func (r *AliexpresssocialitemsearchAPIRequest) SetCommissionRateMax(_commissionRateMax string) error {
	r._commissionRateMax = _commissionRateMax
	r.Set("commission_rate_max", _commissionRateMax)
	return nil
}

// GetCommissionRateMax CommissionRateMax Getter
func (r AliexpresssocialitemsearchAPIRequest) GetCommissionRateMax() string {
	return r._commissionRateMax
}

// SetCommissionRateMin is CommissionRateMin Setter
// 佣金最小值
func (r *AliexpresssocialitemsearchAPIRequest) SetCommissionRateMin(_commissionRateMin string) error {
	r._commissionRateMin = _commissionRateMin
	r.Set("commission_rate_min", _commissionRateMin)
	return nil
}

// GetCommissionRateMin CommissionRateMin Getter
func (r AliexpresssocialitemsearchAPIRequest) GetCommissionRateMin() string {
	return r._commissionRateMin
}

// SetMinPrice is MinPrice Setter
// 最低价格
func (r *AliexpresssocialitemsearchAPIRequest) SetMinPrice(_minPrice string) error {
	r._minPrice = _minPrice
	r.Set("min_price", _minPrice)
	return nil
}

// GetMinPrice MinPrice Getter
func (r AliexpresssocialitemsearchAPIRequest) GetMinPrice() string {
	return r._minPrice
}

// SetMaxPrice is MaxPrice Setter
// 最高价格
func (r *AliexpresssocialitemsearchAPIRequest) SetMaxPrice(_maxPrice string) error {
	r._maxPrice = _maxPrice
	r.Set("max_price", _maxPrice)
	return nil
}

// GetMaxPrice MaxPrice Getter
func (r AliexpresssocialitemsearchAPIRequest) GetMaxPrice() string {
	return r._maxPrice
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *AliexpresssocialitemsearchAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AliexpresssocialitemsearchAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetShipTo is ShipTo Setter
// shipTo国家
func (r *AliexpresssocialitemsearchAPIRequest) SetShipTo(_shipTo string) error {
	r._shipTo = _shipTo
	r.Set("ship_to", _shipTo)
	return nil
}

// GetShipTo ShipTo Getter
func (r AliexpresssocialitemsearchAPIRequest) GetShipTo() string {
	return r._shipTo
}

// SetCommentScore is CommentScore Setter
// 评价分数
func (r *AliexpresssocialitemsearchAPIRequest) SetCommentScore(_commentScore string) error {
	r._commentScore = _commentScore
	r.Set("comment_score", _commentScore)
	return nil
}

// GetCommentScore CommentScore Getter
func (r AliexpresssocialitemsearchAPIRequest) GetCommentScore() string {
	return r._commentScore
}

// SetCurrency is Currency Setter
// 币种
func (r *AliexpresssocialitemsearchAPIRequest) SetCurrency(_currency string) error {
	r._currency = _currency
	r.Set("currency", _currency)
	return nil
}

// GetCurrency Currency Getter
func (r AliexpresssocialitemsearchAPIRequest) GetCurrency() string {
	return r._currency
}

// SetLocale is Locale Setter
// locale,格式为language+&#34;_&#34;+country
func (r *AliexpresssocialitemsearchAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AliexpresssocialitemsearchAPIRequest) GetLocale() string {
	return r._locale
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpresssocialitemsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpresssocialitemsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetLogisticsTime is LogisticsTime Setter
// 物流时效
func (r *AliexpresssocialitemsearchAPIRequest) SetLogisticsTime(_logisticsTime int64) error {
	r._logisticsTime = _logisticsTime
	r.Set("logistics_time", _logisticsTime)
	return nil
}

// GetLogisticsTime LogisticsTime Getter
func (r AliexpresssocialitemsearchAPIRequest) GetLogisticsTime() int64 {
	return r._logisticsTime
}

// SetCateId is CateId Setter
// 类目ID
func (r *AliexpresssocialitemsearchAPIRequest) SetCateId(_cateId int64) error {
	r._cateId = _cateId
	r.Set("cate_id", _cateId)
	return nil
}

// GetCateId CateId Getter
func (r AliexpresssocialitemsearchAPIRequest) GetCateId() int64 {
	return r._cateId
}

// SetPageNo is PageNo Setter
// 页码
func (r *AliexpresssocialitemsearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpresssocialitemsearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetHasVideo is HasVideo Setter
// 是否有视频
func (r *AliexpresssocialitemsearchAPIRequest) SetHasVideo(_hasVideo bool) error {
	r._hasVideo = _hasVideo
	r.Set("has_video", _hasVideo)
	return nil
}

// GetHasVideo HasVideo Getter
func (r AliexpresssocialitemsearchAPIRequest) GetHasVideo() bool {
	return r._hasVideo
}

// SetDesc is Desc Setter
// 是否逆序
func (r *AliexpresssocialitemsearchAPIRequest) SetDesc(_desc bool) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r AliexpresssocialitemsearchAPIRequest) GetDesc() bool {
	return r._desc
}

// SetIsShipFree is IsShipFree Setter
// 是否免邮
func (r *AliexpresssocialitemsearchAPIRequest) SetIsShipFree(_isShipFree bool) error {
	r._isShipFree = _isShipFree
	r.Set("is_ship_free", _isShipFree)
	return nil
}

// GetIsShipFree IsShipFree Getter
func (r AliexpresssocialitemsearchAPIRequest) GetIsShipFree() bool {
	return r._isShipFree
}

// SetAllianceItem is AllianceItem Setter
// 是否联盟商品
func (r *AliexpresssocialitemsearchAPIRequest) SetAllianceItem(_allianceItem bool) error {
	r._allianceItem = _allianceItem
	r.Set("alliance_item", _allianceItem)
	return nil
}

// GetAllianceItem AllianceItem Getter
func (r AliexpresssocialitemsearchAPIRequest) GetAllianceItem() bool {
	return r._allianceItem
}
