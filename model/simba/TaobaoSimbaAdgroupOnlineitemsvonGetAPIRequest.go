package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest 获取用户上架在线销售的全部宝贝 API请求
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 排序字段，starts：按开始时间排序bidCount:按销量排序
	_orderField string
	// 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
	_productId int64
	// 页尺寸，最大200
	_pageSize int64
	// 页码，从1开始,最大50。最大只能获取1W个宝贝
	_pageNo int64
	// 排序，true:降序， false:升序
	_orderBy bool
}

// NewTaobaoSimbaAdgroupOnlineitemsvonGetRequest 初始化TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest对象
func NewTaobaoSimbaAdgroupOnlineitemsvonGetRequest() *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest {
	return &TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.onlineitemsvon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetNick() string {
	return r._nick
}

// SetOrderField is OrderField Setter
// 排序字段，starts：按开始时间排序bidCount:按销量排序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetOrderField(_orderField string) error {
	r._orderField = _orderField
	r.Set("order_field", _orderField)
	return nil
}

// GetOrderField OrderField Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetOrderField() string {
	return r._orderField
}

// SetProductId is ProductId Setter
// 推广单元类型 101001005代表标准推广，101001014代表销量明星推广
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPageSize is PageSize Setter
// 页尺寸，最大200
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，从1开始,最大50。最大只能获取1W个宝贝
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetOrderBy is OrderBy Setter
// 排序，true:降序， false:升序
func (r *TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) SetOrderBy(_orderBy bool) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest) GetOrderBy() bool {
	return r._orderBy
}
