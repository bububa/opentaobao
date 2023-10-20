package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemsinventorygetAPIRequest 得到当前会话用户库存中的商品列表 API请求
// taobao.items.inventory.get
//
// 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取&lt;br/&gt;
// &lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoitemsinventorygetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值为 Item 商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。
	_fields string
	// 搜索字段。搜索商品的title。
	_q string
	// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>)
	_sellerCids string
	// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
	_orderBy string
	// 分类字段。可选值:<br>regular_shelved(定时上架)<br>never_on_shelf(从未上架)<br>off_shelf(我下架的)<br><font color='red'>for_shelved(等待所有上架)<br>sold_out(全部卖完)<br>violation_off_shelf(违规下架的)<br>默认查询for_shelved(等待所有上架)这个状态的商品<br></font>注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)
	_banner string
	// 商品起始修改时间
	_startModified string
	// 商品结束修改时间
	_endModified string
	// 商品类型：a-拍卖,b-一口价
	_auctionType string
	// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
	_cid int64
	// 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。
	_pageNo int64
	// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。
	_pageSize int64
	// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
	_hasDiscount bool
	// 商品是否在淘宝显示
	_isTaobao bool
	// 商品是否在外部网店显示
	_isEx bool
}

// NewTaobaoitemsinventorygetRequest 初始化TaobaoitemsinventorygetAPIRequest对象
func NewTaobaoitemsinventorygetRequest() *TaobaoitemsinventorygetAPIRequest {
	return &TaobaoitemsinventorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemsinventorygetAPIRequest) GetApiMethodName() string {
	return "taobao.items.inventory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemsinventorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemsinventorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值为 Item 商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。&lt;br&gt; 不支持其他字段，如果需要获取其他字段数据，调用taobao.item.seller.get。
func (r *TaobaoitemsinventorygetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoitemsinventorygetAPIRequest) GetFields() string {
	return r._fields
}

// SetQ is Q Setter
// 搜索字段。搜索商品的title。
func (r *TaobaoitemsinventorygetAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoitemsinventorygetAPIRequest) GetQ() string {
	return r._q
}

// SetSellerCids is SellerCids Setter
// 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(&lt;font color=&#34;red&#34;&gt;注：目前最多支持32个ID号传入&lt;/font&gt;)
func (r *TaobaoitemsinventorygetAPIRequest) SetSellerCids(_sellerCids string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r TaobaoitemsinventorygetAPIRequest) GetSellerCids() string {
	return r._sellerCids
}

// SetOrderBy is OrderBy Setter
// 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
func (r *TaobaoitemsinventorygetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoitemsinventorygetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetBanner is Banner Setter
// 分类字段。可选值:&lt;br&gt;regular_shelved(定时上架)&lt;br&gt;never_on_shelf(从未上架)&lt;br&gt;off_shelf(我下架的)&lt;br&gt;&lt;font color=&#39;red&#39;&gt;for_shelved(等待所有上架)&lt;br&gt;sold_out(全部卖完)&lt;br&gt;violation_off_shelf(违规下架的)&lt;br&gt;默认查询for_shelved(等待所有上架)这个状态的商品&lt;br&gt;&lt;/font&gt;注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)
func (r *TaobaoitemsinventorygetAPIRequest) SetBanner(_banner string) error {
	r._banner = _banner
	r.Set("banner", _banner)
	return nil
}

// GetBanner Banner Getter
func (r TaobaoitemsinventorygetAPIRequest) GetBanner() string {
	return r._banner
}

// SetStartModified is StartModified Setter
// 商品起始修改时间
func (r *TaobaoitemsinventorygetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaoitemsinventorygetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetEndModified is EndModified Setter
// 商品结束修改时间
func (r *TaobaoitemsinventorygetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaoitemsinventorygetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetAuctionType is AuctionType Setter
// 商品类型：a-拍卖,b-一口价
func (r *TaobaoitemsinventorygetAPIRequest) SetAuctionType(_auctionType string) error {
	r._auctionType = _auctionType
	r.Set("auction_type", _auctionType)
	return nil
}

// GetAuctionType AuctionType Getter
func (r TaobaoitemsinventorygetAPIRequest) GetAuctionType() string {
	return r._auctionType
}

// SetCid is Cid Setter
// 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (r *TaobaoitemsinventorygetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoitemsinventorygetAPIRequest) GetCid() int64 {
	return r._cid
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。
func (r *TaobaoitemsinventorygetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoitemsinventorygetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。
func (r *TaobaoitemsinventorygetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoitemsinventorygetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetHasDiscount is HasDiscount Setter
// 是否参与会员折扣。可选值：true，false。默认不过滤该条件
func (r *TaobaoitemsinventorygetAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r TaobaoitemsinventorygetAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}

// SetIsTaobao is IsTaobao Setter
// 商品是否在淘宝显示
func (r *TaobaoitemsinventorygetAPIRequest) SetIsTaobao(_isTaobao bool) error {
	r._isTaobao = _isTaobao
	r.Set("is_taobao", _isTaobao)
	return nil
}

// GetIsTaobao IsTaobao Getter
func (r TaobaoitemsinventorygetAPIRequest) GetIsTaobao() bool {
	return r._isTaobao
}

// SetIsEx is IsEx Setter
// 商品是否在外部网店显示
func (r *TaobaoitemsinventorygetAPIRequest) SetIsEx(_isEx bool) error {
	r._isEx = _isEx
	r.Set("is_ex", _isEx)
	return nil
}

// GetIsEx IsEx Getter
func (r TaobaoitemsinventorygetAPIRequest) GetIsEx() bool {
	return r._isEx
}
