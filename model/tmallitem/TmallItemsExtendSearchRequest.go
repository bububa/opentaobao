package tmallitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索天猫商品 API请求
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
*/
type TmallItemsExtendSearchRequest struct {
    model.Params
    // 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
    _q   string
    // 前台类目id，支持多选过滤，cat=catid1,catid2
    _cat   string
    // 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
    _sort   string
    // 品牌的id。支持多选过滤，brand=brand1,brand2
    _brand   string
    // 宝贝卖家所在地，中文gbk编码
    _loc   string
    // 以“属性id：属性值”的形式传入;
    _prop   string
    // 是否包邮，-1为包邮
    _postFee   int64
    // 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
    _startPrice   float64
    // 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
    _endPrice   float64
    // 是否货到付款，1为货到付款
    _supportCod   int64
    // 是否多倍积分，1为多倍积分
    _manyPoints   int64
    // 显示旺旺在线卖家的宝贝时，wwonline=1
    _wwonline   int64
    // 过滤vip宝贝时，vip=1
    _vip   int64
    // 过滤搭配减价宝贝时，combo=1
    _combo   int64
    // 过滤折扣宝贝时，miaosha=1
    _miaosha   int64
    // 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
    _nspu   int64
    // 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
    _auctionTag   string
    // 可以根据产品Id搜索属于这个spu的商品。
    _spuid   int64
    // 可以根据卖家id搜索属于该卖家的商品
    _userId   int64
    // 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
    _pageNo   int64
    // 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
    _pageSize   int64
    // 后台类目id，category=categoryId
    _category   string
}

// 初始化TmallItemsExtendSearchRequest对象
func NewTmallItemsExtendSearchRequest() *TmallItemsExtendSearchRequest{
    return &TmallItemsExtendSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemsExtendSearchRequest) GetApiMethodName() string {
    return "tmall.items.extend.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemsExtendSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Q Setter
// 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
func (r *TmallItemsExtendSearchRequest) SetQ(_q string) error {
    r._q = _q
    r.Set("q", _q)
    return nil
}

// Q Getter
func (r TmallItemsExtendSearchRequest) GetQ() string {
    return r._q
}
// Cat Setter
// 前台类目id，支持多选过滤，cat=catid1,catid2
func (r *TmallItemsExtendSearchRequest) SetCat(_cat string) error {
    r._cat = _cat
    r.Set("cat", _cat)
    return nil
}

// Cat Getter
func (r TmallItemsExtendSearchRequest) GetCat() string {
    return r._cat
}
// Sort Setter
// 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
func (r *TmallItemsExtendSearchRequest) SetSort(_sort string) error {
    r._sort = _sort
    r.Set("sort", _sort)
    return nil
}

// Sort Getter
func (r TmallItemsExtendSearchRequest) GetSort() string {
    return r._sort
}
// Brand Setter
// 品牌的id。支持多选过滤，brand=brand1,brand2
func (r *TmallItemsExtendSearchRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TmallItemsExtendSearchRequest) GetBrand() string {
    return r._brand
}
// Loc Setter
// 宝贝卖家所在地，中文gbk编码
func (r *TmallItemsExtendSearchRequest) SetLoc(_loc string) error {
    r._loc = _loc
    r.Set("loc", _loc)
    return nil
}

// Loc Getter
func (r TmallItemsExtendSearchRequest) GetLoc() string {
    return r._loc
}
// Prop Setter
// 以“属性id：属性值”的形式传入;
func (r *TmallItemsExtendSearchRequest) SetProp(_prop string) error {
    r._prop = _prop
    r.Set("prop", _prop)
    return nil
}

// Prop Getter
func (r TmallItemsExtendSearchRequest) GetProp() string {
    return r._prop
}
// PostFee Setter
// 是否包邮，-1为包邮
func (r *TmallItemsExtendSearchRequest) SetPostFee(_postFee int64) error {
    r._postFee = _postFee
    r.Set("post_fee", _postFee)
    return nil
}

// PostFee Getter
func (r TmallItemsExtendSearchRequest) GetPostFee() int64 {
    return r._postFee
}
// StartPrice Setter
// 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
func (r *TmallItemsExtendSearchRequest) SetStartPrice(_startPrice float64) error {
    r._startPrice = _startPrice
    r.Set("start_price", _startPrice)
    return nil
}

// StartPrice Getter
func (r TmallItemsExtendSearchRequest) GetStartPrice() float64 {
    return r._startPrice
}
// EndPrice Setter
// 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
func (r *TmallItemsExtendSearchRequest) SetEndPrice(_endPrice float64) error {
    r._endPrice = _endPrice
    r.Set("end_price", _endPrice)
    return nil
}

// EndPrice Getter
func (r TmallItemsExtendSearchRequest) GetEndPrice() float64 {
    return r._endPrice
}
// SupportCod Setter
// 是否货到付款，1为货到付款
func (r *TmallItemsExtendSearchRequest) SetSupportCod(_supportCod int64) error {
    r._supportCod = _supportCod
    r.Set("support_cod", _supportCod)
    return nil
}

// SupportCod Getter
func (r TmallItemsExtendSearchRequest) GetSupportCod() int64 {
    return r._supportCod
}
// ManyPoints Setter
// 是否多倍积分，1为多倍积分
func (r *TmallItemsExtendSearchRequest) SetManyPoints(_manyPoints int64) error {
    r._manyPoints = _manyPoints
    r.Set("many_points", _manyPoints)
    return nil
}

// ManyPoints Getter
func (r TmallItemsExtendSearchRequest) GetManyPoints() int64 {
    return r._manyPoints
}
// Wwonline Setter
// 显示旺旺在线卖家的宝贝时，wwonline=1
func (r *TmallItemsExtendSearchRequest) SetWwonline(_wwonline int64) error {
    r._wwonline = _wwonline
    r.Set("wwonline", _wwonline)
    return nil
}

// Wwonline Getter
func (r TmallItemsExtendSearchRequest) GetWwonline() int64 {
    return r._wwonline
}
// Vip Setter
// 过滤vip宝贝时，vip=1
func (r *TmallItemsExtendSearchRequest) SetVip(_vip int64) error {
    r._vip = _vip
    r.Set("vip", _vip)
    return nil
}

// Vip Getter
func (r TmallItemsExtendSearchRequest) GetVip() int64 {
    return r._vip
}
// Combo Setter
// 过滤搭配减价宝贝时，combo=1
func (r *TmallItemsExtendSearchRequest) SetCombo(_combo int64) error {
    r._combo = _combo
    r.Set("combo", _combo)
    return nil
}

// Combo Getter
func (r TmallItemsExtendSearchRequest) GetCombo() int64 {
    return r._combo
}
// Miaosha Setter
// 过滤折扣宝贝时，miaosha=1
func (r *TmallItemsExtendSearchRequest) SetMiaosha(_miaosha int64) error {
    r._miaosha = _miaosha
    r.Set("miaosha", _miaosha)
    return nil
}

// Miaosha Getter
func (r TmallItemsExtendSearchRequest) GetMiaosha() int64 {
    return r._miaosha
}
// Nspu Setter
// 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
func (r *TmallItemsExtendSearchRequest) SetNspu(_nspu int64) error {
    r._nspu = _nspu
    r.Set("nspu", _nspu)
    return nil
}

// Nspu Getter
func (r TmallItemsExtendSearchRequest) GetNspu() int64 {
    return r._nspu
}
// AuctionTag Setter
// 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
func (r *TmallItemsExtendSearchRequest) SetAuctionTag(_auctionTag string) error {
    r._auctionTag = _auctionTag
    r.Set("auction_tag", _auctionTag)
    return nil
}

// AuctionTag Getter
func (r TmallItemsExtendSearchRequest) GetAuctionTag() string {
    return r._auctionTag
}
// Spuid Setter
// 可以根据产品Id搜索属于这个spu的商品。
func (r *TmallItemsExtendSearchRequest) SetSpuid(_spuid int64) error {
    r._spuid = _spuid
    r.Set("spuid", _spuid)
    return nil
}

// Spuid Getter
func (r TmallItemsExtendSearchRequest) GetSpuid() int64 {
    return r._spuid
}
// UserId Setter
// 可以根据卖家id搜索属于该卖家的商品
func (r *TmallItemsExtendSearchRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TmallItemsExtendSearchRequest) GetUserId() int64 {
    return r._userId
}
// PageNo Setter
// 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
func (r *TmallItemsExtendSearchRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TmallItemsExtendSearchRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
func (r *TmallItemsExtendSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallItemsExtendSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// Category Setter
// 后台类目id，category=categoryId
func (r *TmallItemsExtendSearchRequest) SetCategory(_category string) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r TmallItemsExtendSearchRequest) GetCategory() string {
    return r._category
}
