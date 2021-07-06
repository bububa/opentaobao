package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemsOnsaleGetAPIRequest 获取当前会话用户出售中的商品列表 API请求
// taobao.items.onsale.get
//
// 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
type TaobaoItemsOnsaleGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
	_fields string
	// 搜索字段。搜索商品的title。
	_q string
	// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
	_sellerCids string
	// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
	_orderBy string
	// 起始的修改时间
	_startModified string
	// 结束的修改时间
	_endModified string
	// 商品类型：a-拍卖,b-一口价
	_auctionType string
	// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
	_cid int64
	// 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
	_pageNo int64
	// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
	_pageSize int64
	// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
	_hasDiscount bool
	// 是否橱窗推荐。 可选值：true，false。默认不过滤该条件
	_hasShowcase bool
	// 商品是否在淘宝显示
	_isTaobao bool
	// 商品是否在外部网店显示
	_isEx bool
	// 是否挂接了达尔文标准产品体系
	_isCspu bool
	// 组合商品
	_isCombine bool
}

// NewTaobaoItemsOnsaleGetRequest 初始化TaobaoItemsOnsaleGetAPIRequest对象
func NewTaobaoItemsOnsaleGetRequest() *TaobaoItemsOnsaleGetAPIRequest {
	return &TaobaoItemsOnsaleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemsOnsaleGetAPIRequest) GetApiMethodName() string {
	return "taobao.items.onsale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemsOnsaleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id,sold_quantity ；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get 获取。
func (r *TaobaoItemsOnsaleGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetFields() string {
	return r._fields
}

// SetQ is Q Setter
// 搜索字段。搜索商品的title。
func (r *TaobaoItemsOnsaleGetAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetQ() string {
	return r._q
}

// SetSellerCids is SellerCids Setter
// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
func (r *TaobaoItemsOnsaleGetAPIRequest) SetSellerCids(_sellerCids string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetSellerCids() string {
	return r._sellerCids
}

// SetOrderBy is OrderBy Setter
// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
func (r *TaobaoItemsOnsaleGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetStartModified is StartModified Setter
// 起始的修改时间
func (r *TaobaoItemsOnsaleGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetEndModified is EndModified Setter
// 结束的修改时间
func (r *TaobaoItemsOnsaleGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetAuctionType is AuctionType Setter
// 商品类型：a-拍卖,b-一口价
func (r *TaobaoItemsOnsaleGetAPIRequest) SetAuctionType(_auctionType string) error {
	r._auctionType = _auctionType
	r.Set("auction_type", _auctionType)
	return nil
}

// GetAuctionType AuctionType Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetAuctionType() string {
	return r._auctionType
}

// SetCid is Cid Setter
// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (r *TaobaoItemsOnsaleGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetCid() int64 {
	return r._cid
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (r *TaobaoItemsOnsaleGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (r *TaobaoItemsOnsaleGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetHasDiscount is HasDiscount Setter
// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
func (r *TaobaoItemsOnsaleGetAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}

// SetHasShowcase is HasShowcase Setter
// 是否橱窗推荐。 可选值：true，false。默认不过滤该条件
func (r *TaobaoItemsOnsaleGetAPIRequest) SetHasShowcase(_hasShowcase bool) error {
	r._hasShowcase = _hasShowcase
	r.Set("has_showcase", _hasShowcase)
	return nil
}

// GetHasShowcase HasShowcase Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetHasShowcase() bool {
	return r._hasShowcase
}

// SetIsTaobao is IsTaobao Setter
// 商品是否在淘宝显示
func (r *TaobaoItemsOnsaleGetAPIRequest) SetIsTaobao(_isTaobao bool) error {
	r._isTaobao = _isTaobao
	r.Set("is_taobao", _isTaobao)
	return nil
}

// GetIsTaobao IsTaobao Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetIsTaobao() bool {
	return r._isTaobao
}

// SetIsEx is IsEx Setter
// 商品是否在外部网店显示
func (r *TaobaoItemsOnsaleGetAPIRequest) SetIsEx(_isEx bool) error {
	r._isEx = _isEx
	r.Set("is_ex", _isEx)
	return nil
}

// GetIsEx IsEx Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetIsEx() bool {
	return r._isEx
}

// SetIsCspu is IsCspu Setter
// 是否挂接了达尔文标准产品体系
func (r *TaobaoItemsOnsaleGetAPIRequest) SetIsCspu(_isCspu bool) error {
	r._isCspu = _isCspu
	r.Set("is_cspu", _isCspu)
	return nil
}

// GetIsCspu IsCspu Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetIsCspu() bool {
	return r._isCspu
}

// SetIsCombine is IsCombine Setter
// 组合商品
func (r *TaobaoItemsOnsaleGetAPIRequest) SetIsCombine(_isCombine bool) error {
	r._isCombine = _isCombine
	r.Set("is_combine", _isCombine)
	return nil
}

// GetIsCombine IsCombine Getter
func (r TaobaoItemsOnsaleGetAPIRequest) GetIsCombine() bool {
	return r._isCombine
}
