package tmallitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsextendsearchAPIRequest 搜索天猫商品 API请求
// tmall.items.extend.search
//
// 提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
type TmallitemsextendsearchAPIRequest struct {
	model.Params
	// 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
	_auctionTag string
	// 品牌的id。支持多选过滤，brand=brand1,brand2
	_brand string
	// 前台类目id，支持多选过滤，cat=catid1,catid2
	_cat string
	// 后台类目id，category=categoryId
	_category string
	// 宝贝卖家所在地，中文gbk编码
	_loc string
	// 以“属性id：属性值”的形式传入;
	_prop string
	// 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
	_q string
	// 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
	_sort string
	// 过滤搭配减价宝贝时，combo=1
	_combo int64
	// 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
	_endPrice float64
	// 是否多倍积分，1为多倍积分
	_manyPoints int64
	// 过滤折扣宝贝时，miaosha=1
	_miaosha int64
	// 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
	_nspu int64
	// 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
	_pageNo int64
	// 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
	_pageSize int64
	// 是否包邮，-1为包邮
	_postFee int64
	// 可以根据产品Id搜索属于这个spu的商品。
	_spuid int64
	// 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
	_startPrice float64
	// 是否货到付款，1为货到付款
	_supportCod int64
	// 可以根据卖家id搜索属于该卖家的商品
	_userId int64
	// 过滤vip宝贝时，vip=1
	_vip int64
	// 显示旺旺在线卖家的宝贝时，wwonline=1
	_wwonline int64
}

// NewTmallitemsextendsearchRequest 初始化TmallitemsextendsearchAPIRequest对象
func NewTmallitemsextendsearchRequest() *TmallitemsextendsearchAPIRequest {
	return &TmallitemsextendsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsextendsearchAPIRequest) GetApiMethodName() string {
	return "tmall.items.extend.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsextendsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsextendsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuctionTag is AuctionTag Setter
// 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
func (r *TmallitemsextendsearchAPIRequest) SetAuctionTag(_auctionTag string) error {
	r._auctionTag = _auctionTag
	r.Set("auction_tag", _auctionTag)
	return nil
}

// GetAuctionTag AuctionTag Getter
func (r TmallitemsextendsearchAPIRequest) GetAuctionTag() string {
	return r._auctionTag
}

// SetBrand is Brand Setter
// 品牌的id。支持多选过滤，brand=brand1,brand2
func (r *TmallitemsextendsearchAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TmallitemsextendsearchAPIRequest) GetBrand() string {
	return r._brand
}

// SetCat is Cat Setter
// 前台类目id，支持多选过滤，cat=catid1,catid2
func (r *TmallitemsextendsearchAPIRequest) SetCat(_cat string) error {
	r._cat = _cat
	r.Set("cat", _cat)
	return nil
}

// GetCat Cat Getter
func (r TmallitemsextendsearchAPIRequest) GetCat() string {
	return r._cat
}

// SetCategory is Category Setter
// 后台类目id，category=categoryId
func (r *TmallitemsextendsearchAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r TmallitemsextendsearchAPIRequest) GetCategory() string {
	return r._category
}

// SetLoc is Loc Setter
// 宝贝卖家所在地，中文gbk编码
func (r *TmallitemsextendsearchAPIRequest) SetLoc(_loc string) error {
	r._loc = _loc
	r.Set("loc", _loc)
	return nil
}

// GetLoc Loc Getter
func (r TmallitemsextendsearchAPIRequest) GetLoc() string {
	return r._loc
}

// SetProp is Prop Setter
// 以“属性id：属性值”的形式传入;
func (r *TmallitemsextendsearchAPIRequest) SetProp(_prop string) error {
	r._prop = _prop
	r.Set("prop", _prop)
	return nil
}

// GetProp Prop Getter
func (r TmallitemsextendsearchAPIRequest) GetProp() string {
	return r._prop
}

// SetQ is Q Setter
// 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
func (r *TmallitemsextendsearchAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TmallitemsextendsearchAPIRequest) GetQ() string {
	return r._q
}

// SetSort is Sort Setter
// 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
func (r *TmallitemsextendsearchAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TmallitemsextendsearchAPIRequest) GetSort() string {
	return r._sort
}

// SetCombo is Combo Setter
// 过滤搭配减价宝贝时，combo=1
func (r *TmallitemsextendsearchAPIRequest) SetCombo(_combo int64) error {
	r._combo = _combo
	r.Set("combo", _combo)
	return nil
}

// GetCombo Combo Getter
func (r TmallitemsextendsearchAPIRequest) GetCombo() int64 {
	return r._combo
}

// SetEndPrice is EndPrice Setter
// 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
func (r *TmallitemsextendsearchAPIRequest) SetEndPrice(_endPrice float64) error {
	r._endPrice = _endPrice
	r.Set("end_price", _endPrice)
	return nil
}

// GetEndPrice EndPrice Getter
func (r TmallitemsextendsearchAPIRequest) GetEndPrice() float64 {
	return r._endPrice
}

// SetManyPoints is ManyPoints Setter
// 是否多倍积分，1为多倍积分
func (r *TmallitemsextendsearchAPIRequest) SetManyPoints(_manyPoints int64) error {
	r._manyPoints = _manyPoints
	r.Set("many_points", _manyPoints)
	return nil
}

// GetManyPoints ManyPoints Getter
func (r TmallitemsextendsearchAPIRequest) GetManyPoints() int64 {
	return r._manyPoints
}

// SetMiaosha is Miaosha Setter
// 过滤折扣宝贝时，miaosha=1
func (r *TmallitemsextendsearchAPIRequest) SetMiaosha(_miaosha int64) error {
	r._miaosha = _miaosha
	r.Set("miaosha", _miaosha)
	return nil
}

// GetMiaosha Miaosha Getter
func (r TmallitemsextendsearchAPIRequest) GetMiaosha() int64 {
	return r._miaosha
}

// SetNspu is Nspu Setter
// 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
func (r *TmallitemsextendsearchAPIRequest) SetNspu(_nspu int64) error {
	r._nspu = _nspu
	r.Set("nspu", _nspu)
	return nil
}

// GetNspu Nspu Getter
func (r TmallitemsextendsearchAPIRequest) GetNspu() int64 {
	return r._nspu
}

// SetPageNo is PageNo Setter
// 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
func (r *TmallitemsextendsearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TmallitemsextendsearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
func (r *TmallitemsextendsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallitemsextendsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPostFee is PostFee Setter
// 是否包邮，-1为包邮
func (r *TmallitemsextendsearchAPIRequest) SetPostFee(_postFee int64) error {
	r._postFee = _postFee
	r.Set("post_fee", _postFee)
	return nil
}

// GetPostFee PostFee Getter
func (r TmallitemsextendsearchAPIRequest) GetPostFee() int64 {
	return r._postFee
}

// SetSpuid is Spuid Setter
// 可以根据产品Id搜索属于这个spu的商品。
func (r *TmallitemsextendsearchAPIRequest) SetSpuid(_spuid int64) error {
	r._spuid = _spuid
	r.Set("spuid", _spuid)
	return nil
}

// GetSpuid Spuid Getter
func (r TmallitemsextendsearchAPIRequest) GetSpuid() int64 {
	return r._spuid
}

// SetStartPrice is StartPrice Setter
// 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
func (r *TmallitemsextendsearchAPIRequest) SetStartPrice(_startPrice float64) error {
	r._startPrice = _startPrice
	r.Set("start_price", _startPrice)
	return nil
}

// GetStartPrice StartPrice Getter
func (r TmallitemsextendsearchAPIRequest) GetStartPrice() float64 {
	return r._startPrice
}

// SetSupportCod is SupportCod Setter
// 是否货到付款，1为货到付款
func (r *TmallitemsextendsearchAPIRequest) SetSupportCod(_supportCod int64) error {
	r._supportCod = _supportCod
	r.Set("support_cod", _supportCod)
	return nil
}

// GetSupportCod SupportCod Getter
func (r TmallitemsextendsearchAPIRequest) GetSupportCod() int64 {
	return r._supportCod
}

// SetUserId is UserId Setter
// 可以根据卖家id搜索属于该卖家的商品
func (r *TmallitemsextendsearchAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TmallitemsextendsearchAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetVip is Vip Setter
// 过滤vip宝贝时，vip=1
func (r *TmallitemsextendsearchAPIRequest) SetVip(_vip int64) error {
	r._vip = _vip
	r.Set("vip", _vip)
	return nil
}

// GetVip Vip Getter
func (r TmallitemsextendsearchAPIRequest) GetVip() int64 {
	return r._vip
}

// SetWwonline is Wwonline Setter
// 显示旺旺在线卖家的宝贝时，wwonline=1
func (r *TmallitemsextendsearchAPIRequest) SetWwonline(_wwonline int64) error {
	r._wwonline = _wwonline
	r.Set("wwonline", _wwonline)
	return nil
}

// GetWwonline Wwonline Getter
func (r TmallitemsextendsearchAPIRequest) GetWwonline() int64 {
	return r._wwonline
}
