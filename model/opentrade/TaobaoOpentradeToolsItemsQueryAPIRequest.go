package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradetoolsitemsqueryAPIRequest 交易开放获取商品绑定信息 API请求
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
type TaobaoopentradetoolsitemsqueryAPIRequest struct {
	model.Params
	// 交易开放C端小程序ID
	_miniappId int64
	// 起始页
	_pageIndex int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoopentradetoolsitemsqueryRequest 初始化TaobaoopentradetoolsitemsqueryAPIRequest对象
func NewTaobaoopentradetoolsitemsqueryRequest() *TaobaoopentradetoolsitemsqueryAPIRequest {
	return &TaobaoopentradetoolsitemsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniappId is MiniappId Setter
// 交易开放C端小程序ID
func (r *TaobaoopentradetoolsitemsqueryAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}

// SetPageIndex is PageIndex Setter
// 起始页
func (r *TaobaoopentradetoolsitemsqueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoopentradetoolsitemsqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopentradetoolsitemsqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
