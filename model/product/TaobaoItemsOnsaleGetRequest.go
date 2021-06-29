package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当前会话用户出售中的商品列表 API请求
taobao.items.onsale.get

获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsOnsaleGetRequest struct {
    model.Params
    // 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
    fields   string
    // 搜索字段。搜索商品的title。
    q   string
    // 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
    cid   int64
    // 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
    sellerCids   string
    // 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
    pageNo   int64
    // 是否参与会员折扣。可选值：true，false。默认不过滤该条件
    hasDiscount   bool
    // 是否橱窗推荐。 可选值：true，false。默认不过滤该条件
    hasShowcase   bool
    // 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
    orderBy   string
    // 商品是否在淘宝显示
    isTaobao   bool
    // 商品是否在外部网店显示
    isEx   bool
    // 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
    pageSize   int64
    // 起始的修改时间
    startModified   string
    // 结束的修改时间
    endModified   string
    // 是否挂接了达尔文标准产品体系
    isCspu   bool
    // 组合商品
    isCombine   bool
    // 商品类型：a-拍卖,b-一口价
    auctionType   string
}

// 初始化TaobaoItemsOnsaleGetRequest对象
func NewTaobaoItemsOnsaleGetRequest() *TaobaoItemsOnsaleGetRequest{
    return &TaobaoItemsOnsaleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemsOnsaleGetRequest) GetApiMethodName() string {
    return "taobao.items.onsale.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemsOnsaleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
func (r *TaobaoItemsOnsaleGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoItemsOnsaleGetRequest) GetFields() string {
    return r.fields
}
// Q Setter
// 搜索字段。搜索商品的title。
func (r *TaobaoItemsOnsaleGetRequest) SetQ(q string) error {
    r.q = q
    r.Set("q", q)
    return nil
}

// Q Getter
func (r TaobaoItemsOnsaleGetRequest) GetQ() string {
    return r.q
}
// Cid Setter
// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (r *TaobaoItemsOnsaleGetRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoItemsOnsaleGetRequest) GetCid() int64 {
    return r.cid
}
// SellerCids Setter
// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
func (r *TaobaoItemsOnsaleGetRequest) SetSellerCids(sellerCids string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

// SellerCids Getter
func (r TaobaoItemsOnsaleGetRequest) GetSellerCids() string {
    return r.sellerCids
}
// PageNo Setter
// 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (r *TaobaoItemsOnsaleGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoItemsOnsaleGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// HasDiscount Setter
// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
func (r *TaobaoItemsOnsaleGetRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

// HasDiscount Getter
func (r TaobaoItemsOnsaleGetRequest) GetHasDiscount() bool {
    return r.hasDiscount
}
// HasShowcase Setter
// 是否橱窗推荐。 可选值：true，false。默认不过滤该条件
func (r *TaobaoItemsOnsaleGetRequest) SetHasShowcase(hasShowcase bool) error {
    r.hasShowcase = hasShowcase
    r.Set("has_showcase", hasShowcase)
    return nil
}

// HasShowcase Getter
func (r TaobaoItemsOnsaleGetRequest) GetHasShowcase() bool {
    return r.hasShowcase
}
// OrderBy Setter
// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
func (r *TaobaoItemsOnsaleGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoItemsOnsaleGetRequest) GetOrderBy() string {
    return r.orderBy
}
// IsTaobao Setter
// 商品是否在淘宝显示
func (r *TaobaoItemsOnsaleGetRequest) SetIsTaobao(isTaobao bool) error {
    r.isTaobao = isTaobao
    r.Set("is_taobao", isTaobao)
    return nil
}

// IsTaobao Getter
func (r TaobaoItemsOnsaleGetRequest) GetIsTaobao() bool {
    return r.isTaobao
}
// IsEx Setter
// 商品是否在外部网店显示
func (r *TaobaoItemsOnsaleGetRequest) SetIsEx(isEx bool) error {
    r.isEx = isEx
    r.Set("is_ex", isEx)
    return nil
}

// IsEx Getter
func (r TaobaoItemsOnsaleGetRequest) GetIsEx() bool {
    return r.isEx
}
// PageSize Setter
// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (r *TaobaoItemsOnsaleGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoItemsOnsaleGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// StartModified Setter
// 起始的修改时间
func (r *TaobaoItemsOnsaleGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

// StartModified Getter
func (r TaobaoItemsOnsaleGetRequest) GetStartModified() string {
    return r.startModified
}
// EndModified Setter
// 结束的修改时间
func (r *TaobaoItemsOnsaleGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

// EndModified Getter
func (r TaobaoItemsOnsaleGetRequest) GetEndModified() string {
    return r.endModified
}
// IsCspu Setter
// 是否挂接了达尔文标准产品体系
func (r *TaobaoItemsOnsaleGetRequest) SetIsCspu(isCspu bool) error {
    r.isCspu = isCspu
    r.Set("is_cspu", isCspu)
    return nil
}

// IsCspu Getter
func (r TaobaoItemsOnsaleGetRequest) GetIsCspu() bool {
    return r.isCspu
}
// IsCombine Setter
// 组合商品
func (r *TaobaoItemsOnsaleGetRequest) SetIsCombine(isCombine bool) error {
    r.isCombine = isCombine
    r.Set("is_combine", isCombine)
    return nil
}

// IsCombine Getter
func (r TaobaoItemsOnsaleGetRequest) GetIsCombine() bool {
    return r.isCombine
}
// AuctionType Setter
// 商品类型：a-拍卖,b-一口价
func (r *TaobaoItemsOnsaleGetRequest) SetAuctionType(auctionType string) error {
    r.auctionType = auctionType
    r.Set("auction_type", auctionType)
    return nil
}

// AuctionType Getter
func (r TaobaoItemsOnsaleGetRequest) GetAuctionType() string {
    return r.auctionType
}
