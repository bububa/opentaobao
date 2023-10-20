package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemQueryAPIRequest 分页查询商品 API请求
// taobao.wlb.item.query
//
// 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryAPIRequest struct {
	model.Params
	// 商品名称
	_name string
	// 商品前台销售名字
	_title string
	// 商家编码
	_itemCode string
	// 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理
	_isSku string
	// 只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理
	_status string
	// ITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理
	_itemType string
	// 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品
	_parentId int64
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// NewTaobaoWlbItemQueryRequest 初始化TaobaoWlbItemQueryAPIRequest对象
func NewTaobaoWlbItemQueryRequest() *TaobaoWlbItemQueryAPIRequest {
	return &TaobaoWlbItemQueryAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbItemQueryAPIRequest) Reset() {
	r._name = ""
	r._title = ""
	r._itemCode = ""
	r._isSku = ""
	r._status = ""
	r._itemType = ""
	r._parentId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 商品名称
func (r *TaobaoWlbItemQueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoWlbItemQueryAPIRequest) GetName() string {
	return r._name
}

// SetTitle is Title Setter
// 商品前台销售名字
func (r *TaobaoWlbItemQueryAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoWlbItemQueryAPIRequest) GetTitle() string {
	return r._title
}

// SetItemCode is ItemCode Setter
// 商家编码
func (r *TaobaoWlbItemQueryAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaoWlbItemQueryAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetIsSku is IsSku Setter
// 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给&#34;true&#34;,&#34;false&#34;来表示;  若值不在范围内，则按true处理
func (r *TaobaoWlbItemQueryAPIRequest) SetIsSku(_isSku string) error {
	r._isSku = _isSku
	r.Set("is_sku", _isSku)
	return nil
}

// GetIsSku IsSku Getter
func (r TaobaoWlbItemQueryAPIRequest) GetIsSku() string {
	return r._isSku
}

// SetStatus is Status Setter
// 只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理
func (r *TaobaoWlbItemQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoWlbItemQueryAPIRequest) GetStatus() string {
	return r._status
}

// SetItemType is ItemType Setter
// ITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理
func (r *TaobaoWlbItemQueryAPIRequest) SetItemType(_itemType string) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r TaobaoWlbItemQueryAPIRequest) GetItemType() string {
	return r._itemType
}

// SetParentId is ParentId Setter
// 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品
func (r *TaobaoWlbItemQueryAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r TaobaoWlbItemQueryAPIRequest) GetParentId() int64 {
	return r._parentId
}

// SetPageNo is PageNo Setter
// 当前页
func (r *TaobaoWlbItemQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbItemQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoWlbItemQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbItemQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoWlbItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbItemQueryRequest()
	},
}

// GetTaobaoWlbItemQueryRequest 从 sync.Pool 获取 TaobaoWlbItemQueryAPIRequest
func GetTaobaoWlbItemQueryAPIRequest() *TaobaoWlbItemQueryAPIRequest {
	return poolTaobaoWlbItemQueryAPIRequest.Get().(*TaobaoWlbItemQueryAPIRequest)
}

// ReleaseTaobaoWlbItemQueryAPIRequest 将 TaobaoWlbItemQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbItemQueryAPIRequest(v *TaobaoWlbItemQueryAPIRequest) {
	v.Reset()
	poolTaobaoWlbItemQueryAPIRequest.Put(v)
}
