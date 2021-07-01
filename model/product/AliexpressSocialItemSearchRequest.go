package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE社交选品 API请求
aliexpress.social.item.search

AE社交选品,通过各种筛选条件对社交商品池进行筛选
*/
type AliexpressSocialItemSearchAPIRequest struct {
    model.Params
    // 是否有视频
    _hasVideo   bool
    // order by properties
    _orderBy   string
    // 是否逆序
    _desc   bool
    // page size
    _pageSize   int64
    // 是否免邮
    _isShipFree   bool
    // 佣金最大值
    _commissionRateMax   string
    // 佣金最小值
    _commissionRateMin   string
    // 物流时效
    _logisticsTime   int64
    // 是否联盟商品
    _allianceItem   bool
    // 类目ID
    _cateId   int64
    // 页码
    _pageNo   int64
    // 最低价格
    _minPrice   string
    // 最高价格
    _maxPrice   string
    // 搜索关键字
    _keyword   string
    // shipTo国家
    _shipTo   string
    // 评价分数
    _commentScore   string
    // 币种
    _currency   string
    // locale,格式为language+"_"+country
    _locale   string
}

// 初始化AliexpressSocialItemSearchAPIRequest对象
func NewAliexpressSocialItemSearchRequest() *AliexpressSocialItemSearchAPIRequest{
    return &AliexpressSocialItemSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialItemSearchAPIRequest) GetApiMethodName() string {
    return "aliexpress.social.item.search"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialItemSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HasVideo Setter
// 是否有视频
func (r *AliexpressSocialItemSearchAPIRequest) SetHasVideo(_hasVideo bool) error {
    r._hasVideo = _hasVideo
    r.Set("has_video", _hasVideo)
    return nil
}

// HasVideo Getter
func (r AliexpressSocialItemSearchAPIRequest) GetHasVideo() bool {
    return r._hasVideo
}
// OrderBy Setter
// order by properties
func (r *AliexpressSocialItemSearchAPIRequest) SetOrderBy(_orderBy string) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r AliexpressSocialItemSearchAPIRequest) GetOrderBy() string {
    return r._orderBy
}
// Desc Setter
// 是否逆序
func (r *AliexpressSocialItemSearchAPIRequest) SetDesc(_desc bool) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r AliexpressSocialItemSearchAPIRequest) GetDesc() bool {
    return r._desc
}
// PageSize Setter
// page size
func (r *AliexpressSocialItemSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSocialItemSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// IsShipFree Setter
// 是否免邮
func (r *AliexpressSocialItemSearchAPIRequest) SetIsShipFree(_isShipFree bool) error {
    r._isShipFree = _isShipFree
    r.Set("is_ship_free", _isShipFree)
    return nil
}

// IsShipFree Getter
func (r AliexpressSocialItemSearchAPIRequest) GetIsShipFree() bool {
    return r._isShipFree
}
// CommissionRateMax Setter
// 佣金最大值
func (r *AliexpressSocialItemSearchAPIRequest) SetCommissionRateMax(_commissionRateMax string) error {
    r._commissionRateMax = _commissionRateMax
    r.Set("commission_rate_max", _commissionRateMax)
    return nil
}

// CommissionRateMax Getter
func (r AliexpressSocialItemSearchAPIRequest) GetCommissionRateMax() string {
    return r._commissionRateMax
}
// CommissionRateMin Setter
// 佣金最小值
func (r *AliexpressSocialItemSearchAPIRequest) SetCommissionRateMin(_commissionRateMin string) error {
    r._commissionRateMin = _commissionRateMin
    r.Set("commission_rate_min", _commissionRateMin)
    return nil
}

// CommissionRateMin Getter
func (r AliexpressSocialItemSearchAPIRequest) GetCommissionRateMin() string {
    return r._commissionRateMin
}
// LogisticsTime Setter
// 物流时效
func (r *AliexpressSocialItemSearchAPIRequest) SetLogisticsTime(_logisticsTime int64) error {
    r._logisticsTime = _logisticsTime
    r.Set("logistics_time", _logisticsTime)
    return nil
}

// LogisticsTime Getter
func (r AliexpressSocialItemSearchAPIRequest) GetLogisticsTime() int64 {
    return r._logisticsTime
}
// AllianceItem Setter
// 是否联盟商品
func (r *AliexpressSocialItemSearchAPIRequest) SetAllianceItem(_allianceItem bool) error {
    r._allianceItem = _allianceItem
    r.Set("alliance_item", _allianceItem)
    return nil
}

// AllianceItem Getter
func (r AliexpressSocialItemSearchAPIRequest) GetAllianceItem() bool {
    return r._allianceItem
}
// CateId Setter
// 类目ID
func (r *AliexpressSocialItemSearchAPIRequest) SetCateId(_cateId int64) error {
    r._cateId = _cateId
    r.Set("cate_id", _cateId)
    return nil
}

// CateId Getter
func (r AliexpressSocialItemSearchAPIRequest) GetCateId() int64 {
    return r._cateId
}
// PageNo Setter
// 页码
func (r *AliexpressSocialItemSearchAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressSocialItemSearchAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// MinPrice Setter
// 最低价格
func (r *AliexpressSocialItemSearchAPIRequest) SetMinPrice(_minPrice string) error {
    r._minPrice = _minPrice
    r.Set("min_price", _minPrice)
    return nil
}

// MinPrice Getter
func (r AliexpressSocialItemSearchAPIRequest) GetMinPrice() string {
    return r._minPrice
}
// MaxPrice Setter
// 最高价格
func (r *AliexpressSocialItemSearchAPIRequest) SetMaxPrice(_maxPrice string) error {
    r._maxPrice = _maxPrice
    r.Set("max_price", _maxPrice)
    return nil
}

// MaxPrice Getter
func (r AliexpressSocialItemSearchAPIRequest) GetMaxPrice() string {
    return r._maxPrice
}
// Keyword Setter
// 搜索关键字
func (r *AliexpressSocialItemSearchAPIRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AliexpressSocialItemSearchAPIRequest) GetKeyword() string {
    return r._keyword
}
// ShipTo Setter
// shipTo国家
func (r *AliexpressSocialItemSearchAPIRequest) SetShipTo(_shipTo string) error {
    r._shipTo = _shipTo
    r.Set("ship_to", _shipTo)
    return nil
}

// ShipTo Getter
func (r AliexpressSocialItemSearchAPIRequest) GetShipTo() string {
    return r._shipTo
}
// CommentScore Setter
// 评价分数
func (r *AliexpressSocialItemSearchAPIRequest) SetCommentScore(_commentScore string) error {
    r._commentScore = _commentScore
    r.Set("comment_score", _commentScore)
    return nil
}

// CommentScore Getter
func (r AliexpressSocialItemSearchAPIRequest) GetCommentScore() string {
    return r._commentScore
}
// Currency Setter
// 币种
func (r *AliexpressSocialItemSearchAPIRequest) SetCurrency(_currency string) error {
    r._currency = _currency
    r.Set("currency", _currency)
    return nil
}

// Currency Getter
func (r AliexpressSocialItemSearchAPIRequest) GetCurrency() string {
    return r._currency
}
// Locale Setter
// locale,格式为language+"_"+country
func (r *AliexpressSocialItemSearchAPIRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r AliexpressSocialItemSearchAPIRequest) GetLocale() string {
    return r._locale
}
