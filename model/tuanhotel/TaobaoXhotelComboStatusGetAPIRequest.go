package tuanhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboStatusGetAPIRequest 酒店宝贝状态查询 API请求
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
type TaobaoXhotelComboStatusGetAPIRequest struct {
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

// NewTaobaoXhotelComboStatusGetRequest 初始化TaobaoXhotelComboStatusGetAPIRequest对象
func NewTaobaoXhotelComboStatusGetRequest() *TaobaoXhotelComboStatusGetAPIRequest {
	return &TaobaoXhotelComboStatusGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelComboStatusGetAPIRequest) Reset() {
	r._itemTitle = ""
	r._itemId = 0
	r._pageSize = 0
	r._currentPageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelComboStatusGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.combo.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelComboStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelComboStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemTitle is ItemTitle Setter
// 宝贝名称
func (r *TaobaoXhotelComboStatusGetAPIRequest) SetItemTitle(_itemTitle string) error {
	r._itemTitle = _itemTitle
	r.Set("item_title", _itemTitle)
	return nil
}

// GetItemTitle ItemTitle Getter
func (r TaobaoXhotelComboStatusGetAPIRequest) GetItemTitle() string {
	return r._itemTitle
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *TaobaoXhotelComboStatusGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoXhotelComboStatusGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoXhotelComboStatusGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelComboStatusGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPageNo is CurrentPageNo Setter
// 当前页
func (r *TaobaoXhotelComboStatusGetAPIRequest) SetCurrentPageNo(_currentPageNo int64) error {
	r._currentPageNo = _currentPageNo
	r.Set("current_page_no", _currentPageNo)
	return nil
}

// GetCurrentPageNo CurrentPageNo Getter
func (r TaobaoXhotelComboStatusGetAPIRequest) GetCurrentPageNo() int64 {
	return r._currentPageNo
}

var poolTaobaoXhotelComboStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelComboStatusGetRequest()
	},
}

// GetTaobaoXhotelComboStatusGetRequest 从 sync.Pool 获取 TaobaoXhotelComboStatusGetAPIRequest
func GetTaobaoXhotelComboStatusGetAPIRequest() *TaobaoXhotelComboStatusGetAPIRequest {
	return poolTaobaoXhotelComboStatusGetAPIRequest.Get().(*TaobaoXhotelComboStatusGetAPIRequest)
}

// ReleaseTaobaoXhotelComboStatusGetAPIRequest 将 TaobaoXhotelComboStatusGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelComboStatusGetAPIRequest(v *TaobaoXhotelComboStatusGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelComboStatusGetAPIRequest.Put(v)
}
