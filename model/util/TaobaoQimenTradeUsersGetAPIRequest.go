package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUsersGetAPIRequest 获取奇门用户列表 API请求
// taobao.qimen.trade.users.get
//
// 获取已开通奇门订单服务的用户列表
type TaobaoQimenTradeUsersGetAPIRequest struct {
	model.Params
	// 页数
	_pageSize int64
	// 每页的数量
	_pageIndex int64
}

// NewTaobaoQimenTradeUsersGetRequest 初始化TaobaoQimenTradeUsersGetAPIRequest对象
func NewTaobaoQimenTradeUsersGetRequest() *TaobaoQimenTradeUsersGetAPIRequest {
	return &TaobaoQimenTradeUsersGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenTradeUsersGetAPIRequest) Reset() {
	r._pageSize = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUsersGetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUsersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenTradeUsersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageSize is PageSize Setter
// 页数
func (r *TaobaoQimenTradeUsersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoQimenTradeUsersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 每页的数量
func (r *TaobaoQimenTradeUsersGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoQimenTradeUsersGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTaobaoQimenTradeUsersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenTradeUsersGetRequest()
	},
}

// GetTaobaoQimenTradeUsersGetRequest 从 sync.Pool 获取 TaobaoQimenTradeUsersGetAPIRequest
func GetTaobaoQimenTradeUsersGetAPIRequest() *TaobaoQimenTradeUsersGetAPIRequest {
	return poolTaobaoQimenTradeUsersGetAPIRequest.Get().(*TaobaoQimenTradeUsersGetAPIRequest)
}

// ReleaseTaobaoQimenTradeUsersGetAPIRequest 将 TaobaoQimenTradeUsersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenTradeUsersGetAPIRequest(v *TaobaoQimenTradeUsersGetAPIRequest) {
	v.Reset()
	poolTaobaoQimenTradeUsersGetAPIRequest.Put(v)
}
