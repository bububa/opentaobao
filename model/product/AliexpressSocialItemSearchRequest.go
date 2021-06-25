package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE社交选品 APIRequest
aliexpress.social.item.search

AE社交选品,通过各种筛选条件对社交商品池进行筛选
*/
type AliexpressSocialItemSearchRequest struct {
    model.Params

    // 是否有视频
    hasVideo   bool 

    // order by properties
    orderBy   string 

    // 是否逆序
    desc   bool 

    // page size
    pageSize   int64 

    // 是否免邮
    isShipFree   bool 

    // 佣金最大值
    commissionRateMax   string 

    // 佣金最小值
    commissionRateMin   string 

    // 物流时效
    logisticsTime   int64 

    // 是否联盟商品
    allianceItem   bool 

    // 类目ID
    cateId   int64 

    // 页码
    pageNo   int64 

    // 最低价格
    minPrice   string 

    // 最高价格
    maxPrice   string 

    // 搜索关键字
    keyword   string 

    // shipTo国家
    shipTo   string 

    // 评价分数
    commentScore   string 

    // 币种
    currency   string 

    // locale,格式为language+"_"+country
    locale   string 

}

func NewAliexpressSocialItemSearchRequest() *AliexpressSocialItemSearchRequest{
    return &AliexpressSocialItemSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialItemSearchRequest) GetApiMethodName() string {
    return "aliexpress.social.item.search"
}

func (r AliexpressSocialItemSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialItemSearchRequest) SetHasVideo(hasVideo bool) error {
    r.hasVideo = hasVideo
    r.Set("has_video", hasVideo)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetHasVideo() bool {
    return r.hasVideo
}

func (r *AliexpressSocialItemSearchRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetOrderBy() string {
    return r.orderBy
}

func (r *AliexpressSocialItemSearchRequest) SetDesc(desc bool) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetDesc() bool {
    return r.desc
}

func (r *AliexpressSocialItemSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressSocialItemSearchRequest) SetIsShipFree(isShipFree bool) error {
    r.isShipFree = isShipFree
    r.Set("is_ship_free", isShipFree)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetIsShipFree() bool {
    return r.isShipFree
}

func (r *AliexpressSocialItemSearchRequest) SetCommissionRateMax(commissionRateMax string) error {
    r.commissionRateMax = commissionRateMax
    r.Set("commission_rate_max", commissionRateMax)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetCommissionRateMax() string {
    return r.commissionRateMax
}

func (r *AliexpressSocialItemSearchRequest) SetCommissionRateMin(commissionRateMin string) error {
    r.commissionRateMin = commissionRateMin
    r.Set("commission_rate_min", commissionRateMin)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetCommissionRateMin() string {
    return r.commissionRateMin
}

func (r *AliexpressSocialItemSearchRequest) SetLogisticsTime(logisticsTime int64) error {
    r.logisticsTime = logisticsTime
    r.Set("logistics_time", logisticsTime)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetLogisticsTime() int64 {
    return r.logisticsTime
}

func (r *AliexpressSocialItemSearchRequest) SetAllianceItem(allianceItem bool) error {
    r.allianceItem = allianceItem
    r.Set("alliance_item", allianceItem)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetAllianceItem() bool {
    return r.allianceItem
}

func (r *AliexpressSocialItemSearchRequest) SetCateId(cateId int64) error {
    r.cateId = cateId
    r.Set("cate_id", cateId)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetCateId() int64 {
    return r.cateId
}

func (r *AliexpressSocialItemSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressSocialItemSearchRequest) SetMinPrice(minPrice string) error {
    r.minPrice = minPrice
    r.Set("min_price", minPrice)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetMinPrice() string {
    return r.minPrice
}

func (r *AliexpressSocialItemSearchRequest) SetMaxPrice(maxPrice string) error {
    r.maxPrice = maxPrice
    r.Set("max_price", maxPrice)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetMaxPrice() string {
    return r.maxPrice
}

func (r *AliexpressSocialItemSearchRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetKeyword() string {
    return r.keyword
}

func (r *AliexpressSocialItemSearchRequest) SetShipTo(shipTo string) error {
    r.shipTo = shipTo
    r.Set("ship_to", shipTo)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetShipTo() string {
    return r.shipTo
}

func (r *AliexpressSocialItemSearchRequest) SetCommentScore(commentScore string) error {
    r.commentScore = commentScore
    r.Set("comment_score", commentScore)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetCommentScore() string {
    return r.commentScore
}

func (r *AliexpressSocialItemSearchRequest) SetCurrency(currency string) error {
    r.currency = currency
    r.Set("currency", currency)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetCurrency() string {
    return r.currency
}

func (r *AliexpressSocialItemSearchRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r AliexpressSocialItemSearchRequest) GetLocale() string {
    return r.locale
}

