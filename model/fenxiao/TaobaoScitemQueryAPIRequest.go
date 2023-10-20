package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemQueryAPIRequest 查询后端商品 API请求
// taobao.scitem.query
//
// 查询后端商品
type TaobaoScitemQueryAPIRequest struct {
	model.Params
	// 商品名称
	_itemName string
	// 商家给商品的一个编码
	_outerCode string
	// 仓库编码
	_wmsCode string
	// 条形码
	_barCode string
	// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
	_itemType int64
	// 当前页码数
	_pageIndex int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// NewTaobaoScitemQueryRequest 初始化TaobaoScitemQueryAPIRequest对象
func NewTaobaoScitemQueryRequest() *TaobaoScitemQueryAPIRequest {
	return &TaobaoScitemQueryAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoScitemQueryAPIRequest) Reset() {
	r._itemName = ""
	r._outerCode = ""
	r._wmsCode = ""
	r._barCode = ""
	r._itemType = 0
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoScitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemName is ItemName Setter
// 商品名称
func (r *TaobaoScitemQueryAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// GetItemName ItemName Getter
func (r TaobaoScitemQueryAPIRequest) GetItemName() string {
	return r._itemName
}

// SetOuterCode is OuterCode Setter
// 商家给商品的一个编码
func (r *TaobaoScitemQueryAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoScitemQueryAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetWmsCode is WmsCode Setter
// 仓库编码
func (r *TaobaoScitemQueryAPIRequest) SetWmsCode(_wmsCode string) error {
	r._wmsCode = _wmsCode
	r.Set("wms_code", _wmsCode)
	return nil
}

// GetWmsCode WmsCode Getter
func (r TaobaoScitemQueryAPIRequest) GetWmsCode() string {
	return r._wmsCode
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoScitemQueryAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoScitemQueryAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetItemType is ItemType Setter
// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
func (r *TaobaoScitemQueryAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r TaobaoScitemQueryAPIRequest) GetItemType() int64 {
	return r._itemType
}

// SetPageIndex is PageIndex Setter
// 当前页码数
func (r *TaobaoScitemQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoScitemQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoScitemQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoScitemQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoScitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoScitemQueryRequest()
	},
}

// GetTaobaoScitemQueryRequest 从 sync.Pool 获取 TaobaoScitemQueryAPIRequest
func GetTaobaoScitemQueryAPIRequest() *TaobaoScitemQueryAPIRequest {
	return poolTaobaoScitemQueryAPIRequest.Get().(*TaobaoScitemQueryAPIRequest)
}

// ReleaseTaobaoScitemQueryAPIRequest 将 TaobaoScitemQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoScitemQueryAPIRequest(v *TaobaoScitemQueryAPIRequest) {
	v.Reset()
	poolTaobaoScitemQueryAPIRequest.Put(v)
}
