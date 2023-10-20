package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcombostatusgetAPIRequest 酒店宝贝状态查询 API请求
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
type TaobaoxhotelcombostatusgetAPIRequest struct {
	model.Params
	// 宝贝名称
	_itemTitle string
	// 宝贝id
	_itemId int64
	// 页大小
	_pageSize int64
	// 当前页
	_currentPageNo int64
}

// NewTaobaoxhotelcombostatusgetRequest 初始化TaobaoxhotelcombostatusgetAPIRequest对象
func NewTaobaoxhotelcombostatusgetRequest() *TaobaoxhotelcombostatusgetAPIRequest {
	return &TaobaoxhotelcombostatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcombostatusgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcombostatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcombostatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemTitle is ItemTitle Setter
// 宝贝名称
func (r *TaobaoxhotelcombostatusgetAPIRequest) SetItemTitle(_itemTitle string) error {
	r._itemTitle = _itemTitle
	r.Set("item_title", _itemTitle)
	return nil
}

// GetItemTitle ItemTitle Getter
func (r TaobaoxhotelcombostatusgetAPIRequest) GetItemTitle() string {
	return r._itemTitle
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoxhotelcombostatusgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoxhotelcombostatusgetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoxhotelcombostatusgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoxhotelcombostatusgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPageNo is CurrentPageNo Setter
// 当前页
func (r *TaobaoxhotelcombostatusgetAPIRequest) SetCurrentPageNo(_currentPageNo int64) error {
	r._currentPageNo = _currentPageNo
	r.Set("current_page_no", _currentPageNo)
	return nil
}

// GetCurrentPageNo CurrentPageNo Getter
func (r TaobaoxhotelcombostatusgetAPIRequest) GetCurrentPageNo() int64 {
	return r._currentPageNo
}
