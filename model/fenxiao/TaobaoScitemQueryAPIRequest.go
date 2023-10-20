package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemqueryAPIRequest 查询后端商品 API请求
// taobao.scitem.query
//
// 查询后端商品
type TaobaoscitemqueryAPIRequest struct {
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

// NewTaobaoscitemqueryRequest 初始化TaobaoscitemqueryAPIRequest对象
func NewTaobaoscitemqueryRequest() *TaobaoscitemqueryAPIRequest {
	return &TaobaoscitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemName is ItemName Setter
// 商品名称
func (r *TaobaoscitemqueryAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// GetItemName ItemName Getter
func (r TaobaoscitemqueryAPIRequest) GetItemName() string {
	return r._itemName
}

// SetOuterCode is OuterCode Setter
// 商家给商品的一个编码
func (r *TaobaoscitemqueryAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoscitemqueryAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetWmsCode is WmsCode Setter
// 仓库编码
func (r *TaobaoscitemqueryAPIRequest) SetWmsCode(_wmsCode string) error {
	r._wmsCode = _wmsCode
	r.Set("wms_code", _wmsCode)
	return nil
}

// GetWmsCode WmsCode Getter
func (r TaobaoscitemqueryAPIRequest) GetWmsCode() string {
	return r._wmsCode
}

// SetBarCode is BarCode Setter
// 条形码
func (r *TaobaoscitemqueryAPIRequest) SetBarCode(_barCode string) error {
	r._barCode = _barCode
	r.Set("bar_code", _barCode)
	return nil
}

// GetBarCode BarCode Getter
func (r TaobaoscitemqueryAPIRequest) GetBarCode() string {
	return r._barCode
}

// SetItemType is ItemType Setter
// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
func (r *TaobaoscitemqueryAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r TaobaoscitemqueryAPIRequest) GetItemType() int64 {
	return r._itemType
}

// SetPageIndex is PageIndex Setter
// 当前页码数
func (r *TaobaoscitemqueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoscitemqueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoscitemqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoscitemqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
