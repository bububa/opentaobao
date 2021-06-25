package tmallitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索天猫商品 APIRequest
tmall.items.extend.search

提供天猫商品搜索结果，需要调用精选商品，请改为调用：tmall.selected.items.search
*/
type TmallItemsExtendSearchRequest struct {
    model.Params

    // 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。
    q   string 

    // 前台类目id，支持多选过滤，cat=catid1,catid2
    cat   string 

    // 排序类型。类型包括：s: 人气排序p: 价格从低到高;pd: 价格从高到低;d: 月销量从高到低;td: 总销量从高到低;pt: 按发布时间排序.
    sort   string 

    // 品牌的id。支持多选过滤，brand=brand1,brand2
    brand   string 

    // 宝贝卖家所在地，中文gbk编码
    loc   string 

    // 以“属性id：属性值”的形式传入;
    prop   string 

    // 是否包邮，-1为包邮
    postFee   int64 

    // 在宝贝页面中进行价格筛选的时候，如果填写了最低价格，就会显示该字段。
    startPrice   float64 

    // 在宝贝页面中进行价格筛选的时候，如果填写了最高价格，就会显示该字段。
    endPrice   float64 

    // 是否货到付款，1为货到付款
    supportCod   int64 

    // 是否多倍积分，1为多倍积分
    manyPoints   int64 

    // 显示旺旺在线卖家的宝贝时，wwonline=1
    wwonline   int64 

    // 过滤vip宝贝时，vip=1
    vip   int64 

    // 过滤搭配减价宝贝时，combo=1
    combo   int64 

    // 过滤折扣宝贝时，miaosha=1
    miaosha   int64 

    // 是否需要spu聚合的开关:1为关闭，不传表示遵循后端聚合逻辑。默认不作spu聚合。
    nspu   int64 

    // 商品标签。支持多选过滤,auction_tag=auction_tag1,auction_tag2,不支持天猫精品库8578
    auctionTag   string 

    // 可以根据产品Id搜索属于这个spu的商品。
    spuid   int64 

    // 可以根据卖家id搜索属于该卖家的商品
    userId   int64 

    // 页码。取值范围：大于零的整数；最大值：100；默认值：1，即默认返回第一页数据。
    pageNo   int64 

    // 每页条数。取值范围：大于零的整数；最大值：100；默认值：40
    pageSize   int64 

    // 后台类目id，category=categoryId
    category   string 

}

func NewTmallItemsExtendSearchRequest() *TmallItemsExtendSearchRequest{
    return &TmallItemsExtendSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemsExtendSearchRequest) GetApiMethodName() string {
    return "tmall.items.extend.search"
}

func (r TmallItemsExtendSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemsExtendSearchRequest) SetQ(q string) error {
    r.q = q
    r.Set("q", q)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetQ() string {
    return r.q
}

func (r *TmallItemsExtendSearchRequest) SetCat(cat string) error {
    r.cat = cat
    r.Set("cat", cat)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetCat() string {
    return r.cat
}

func (r *TmallItemsExtendSearchRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetSort() string {
    return r.sort
}

func (r *TmallItemsExtendSearchRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetBrand() string {
    return r.brand
}

func (r *TmallItemsExtendSearchRequest) SetLoc(loc string) error {
    r.loc = loc
    r.Set("loc", loc)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetLoc() string {
    return r.loc
}

func (r *TmallItemsExtendSearchRequest) SetProp(prop string) error {
    r.prop = prop
    r.Set("prop", prop)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetProp() string {
    return r.prop
}

func (r *TmallItemsExtendSearchRequest) SetPostFee(postFee int64) error {
    r.postFee = postFee
    r.Set("post_fee", postFee)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetPostFee() int64 {
    return r.postFee
}

func (r *TmallItemsExtendSearchRequest) SetStartPrice(startPrice float64) error {
    r.startPrice = startPrice
    r.Set("start_price", startPrice)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetStartPrice() float64 {
    return r.startPrice
}

func (r *TmallItemsExtendSearchRequest) SetEndPrice(endPrice float64) error {
    r.endPrice = endPrice
    r.Set("end_price", endPrice)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetEndPrice() float64 {
    return r.endPrice
}

func (r *TmallItemsExtendSearchRequest) SetSupportCod(supportCod int64) error {
    r.supportCod = supportCod
    r.Set("support_cod", supportCod)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetSupportCod() int64 {
    return r.supportCod
}

func (r *TmallItemsExtendSearchRequest) SetManyPoints(manyPoints int64) error {
    r.manyPoints = manyPoints
    r.Set("many_points", manyPoints)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetManyPoints() int64 {
    return r.manyPoints
}

func (r *TmallItemsExtendSearchRequest) SetWwonline(wwonline int64) error {
    r.wwonline = wwonline
    r.Set("wwonline", wwonline)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetWwonline() int64 {
    return r.wwonline
}

func (r *TmallItemsExtendSearchRequest) SetVip(vip int64) error {
    r.vip = vip
    r.Set("vip", vip)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetVip() int64 {
    return r.vip
}

func (r *TmallItemsExtendSearchRequest) SetCombo(combo int64) error {
    r.combo = combo
    r.Set("combo", combo)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetCombo() int64 {
    return r.combo
}

func (r *TmallItemsExtendSearchRequest) SetMiaosha(miaosha int64) error {
    r.miaosha = miaosha
    r.Set("miaosha", miaosha)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetMiaosha() int64 {
    return r.miaosha
}

func (r *TmallItemsExtendSearchRequest) SetNspu(nspu int64) error {
    r.nspu = nspu
    r.Set("nspu", nspu)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetNspu() int64 {
    return r.nspu
}

func (r *TmallItemsExtendSearchRequest) SetAuctionTag(auctionTag string) error {
    r.auctionTag = auctionTag
    r.Set("auction_tag", auctionTag)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetAuctionTag() string {
    return r.auctionTag
}

func (r *TmallItemsExtendSearchRequest) SetSpuid(spuid int64) error {
    r.spuid = spuid
    r.Set("spuid", spuid)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetSpuid() int64 {
    return r.spuid
}

func (r *TmallItemsExtendSearchRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetUserId() int64 {
    return r.userId
}

func (r *TmallItemsExtendSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TmallItemsExtendSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallItemsExtendSearchRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r TmallItemsExtendSearchRequest) GetCategory() string {
    return r.category
}

