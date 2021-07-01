package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTradeUsersGetAPIRequest
获取奇门用户列表 API请求
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表 */
type TaobaoQimenTradeUsersGetAPIRequest struct {
	model.Params
	// 每页的数量
	_pageIndex int64
	// 页数
	_pageSize int64
}

// NewTaobaoQimenTradeUsersGetRequest 初始化TaobaoQimenTradeUsersGetAPIRequest对象
func NewTaobaoQimenTradeUsersGetRequest() *TaobaoQimenTradeUsersGetAPIRequest {
	return &TaobaoQimenTradeUsersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUsersGetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUsersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageIndex Setter
// 每页的数量
func (r *TaobaoQimenTradeUsersGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoQimenTradeUsersGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// Set is PageSize Setter
// 页数
func (r *TaobaoQimenTradeUsersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoQimenTradeUsersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
