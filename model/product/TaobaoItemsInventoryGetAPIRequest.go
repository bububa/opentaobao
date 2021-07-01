package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemsInventoryGetAPIRequest
得到当前会话用户库存中的商品列表 API请求
taobao.items.inventory.get

获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取<br/>
<strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemsInventoryGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值为 Item 商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。
	_fields string
	// 搜索字段。搜索商品的title。
	_q string
	// 分类字段。可选值:<br>regular_shelved(定时上架)<br>never_on_shelf(从未上架)<br>off_shelf(我下架的)<br><font color='red'>for_shelved(等待所有上架)<br>sold_out(全部卖完)<br>violation_off_shelf(违规下架的)<br>默认查询for_shelved(等待所有上架)这个状态的商品<br></font>注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)
	_banner string
	// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
	_cid int64
	// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
	_sellerCids string
	// 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。
	_pageNo int64
	// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。
	_pageSize int64
	// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
	_hasDiscount bool
	// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
	_orderBy string
	// 商品是否在淘宝显示
	_isTaobao bool
	// 商品是否在外部网店显示
	_isEx bool
	// 商品起始修改时间
	_startModified string
	// 商品结束修改时间
	_endModified string
	// 商品类型：a-拍卖,b-一口价
	_auctionType string
}

// NewTaobaoItemsInventoryGetRequest 初始化TaobaoItemsInventoryGetAPIRequest对象
func NewTaobaoItemsInventoryGetRequest() *TaobaoItemsInventoryGetAPIRequest {
	return &TaobaoItemsInventoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemsInventoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.items.inventory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemsInventoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Fields Setter
// 需返回的字段列表。可选值为 Item 商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。
func (r *TaobaoItemsInventoryGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetFields() string {
	return r._fields
}

// Set is Q Setter
// 搜索字段。搜索商品的title。
func (r *TaobaoItemsInventoryGetAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// Get Q Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetQ() string {
	return r._q
}

// Set is Banner Setter
// 分类字段。可选值:<br>regular_shelved(定时上架)<br>never_on_shelf(从未上架)<br>off_shelf(我下架的)<br><font color='red'>for_shelved(等待所有上架)<br>sold_out(全部卖完)<br>violation_off_shelf(违规下架的)<br>默认查询for_shelved(等待所有上架)这个状态的商品<br></font>注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)
func (r *TaobaoItemsInventoryGetAPIRequest) SetBanner(_banner string) error {
	r._banner = _banner
	r.Set("banner", _banner)
	return nil
}

// Get Banner Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetBanner() string {
	return r._banner
}

// Set is Cid Setter
// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (r *TaobaoItemsInventoryGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// Get Cid Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetCid() int64 {
	return r._cid
}

// Set is SellerCids Setter
// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
func (r *TaobaoItemsInventoryGetAPIRequest) SetSellerCids(_sellerCids string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// Get SellerCids Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetSellerCids() string {
	return r._sellerCids
}

// Set is PageNo Setter
// 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。
func (r *TaobaoItemsInventoryGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。
func (r *TaobaoItemsInventoryGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is HasDiscount Setter
// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
func (r *TaobaoItemsInventoryGetAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// Get HasDiscount Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}

// Set is OrderBy Setter
// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
func (r *TaobaoItemsInventoryGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// Get OrderBy Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// Set is IsTaobao Setter
// 商品是否在淘宝显示
func (r *TaobaoItemsInventoryGetAPIRequest) SetIsTaobao(_isTaobao bool) error {
	r._isTaobao = _isTaobao
	r.Set("is_taobao", _isTaobao)
	return nil
}

// Get IsTaobao Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetIsTaobao() bool {
	return r._isTaobao
}

// Set is IsEx Setter
// 商品是否在外部网店显示
func (r *TaobaoItemsInventoryGetAPIRequest) SetIsEx(_isEx bool) error {
	r._isEx = _isEx
	r.Set("is_ex", _isEx)
	return nil
}

// Get IsEx Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetIsEx() bool {
	return r._isEx
}

// Set is StartModified Setter
// 商品起始修改时间
func (r *TaobaoItemsInventoryGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// Get StartModified Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// Set is EndModified Setter
// 商品结束修改时间
func (r *TaobaoItemsInventoryGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// Get EndModified Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// Set is AuctionType Setter
// 商品类型：a-拍卖,b-一口价
func (r *TaobaoItemsInventoryGetAPIRequest) SetAuctionType(_auctionType string) error {
	r._auctionType = _auctionType
	r.Set("auction_type", _auctionType)
	return nil
}

// Get AuctionType Getter
func (r TaobaoItemsInventoryGetAPIRequest) GetAuctionType() string {
	return r._auctionType
}
