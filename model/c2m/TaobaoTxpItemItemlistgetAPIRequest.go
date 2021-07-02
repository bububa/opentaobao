package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTxpItemItemlistgetAPIRequest 淘小铺商品接口 API请求
// taobao.txp.item.itemlistget
//
// 淘小铺商品的查询服务。
type TaobaoTxpItemItemlistgetAPIRequest struct {
	model.Params
	// 第几页
	_beginPage int64
	// 每页多少条
	_pageSize int64
}

// NewTaobaoTxpItemItemlistgetRequest 初始化TaobaoTxpItemItemlistgetAPIRequest对象
func NewTaobaoTxpItemItemlistgetRequest() *TaobaoTxpItemItemlistgetAPIRequest {
	return &TaobaoTxpItemItemlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTxpItemItemlistgetAPIRequest) GetApiMethodName() string {
	return "taobao.txp.item.itemlistget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTxpItemItemlistgetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBeginPage is BeginPage Setter
// 第几页
func (r *TaobaoTxpItemItemlistgetAPIRequest) SetBeginPage(_beginPage int64) error {
	r._beginPage = _beginPage
	r.Set("begin_page", _beginPage)
	return nil
}

// GetBeginPage BeginPage Getter
func (r TaobaoTxpItemItemlistgetAPIRequest) GetBeginPage() int64 {
	return r._beginPage
}

// SetPageSize is PageSize Setter
// 每页多少条
func (r *TaobaoTxpItemItemlistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTxpItemItemlistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
