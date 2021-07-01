package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreitemQueryAPIRequest
门店关联商品查询接口 API请求
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表 */
type TaobaoQimenStoreitemQueryAPIRequest struct {
	model.Params
	// 当前页面
	_page int64
	// 线上门店id
	_storeId int64
}

// NewTaobaoQimenStoreitemQueryRequest 初始化TaobaoQimenStoreitemQueryAPIRequest对象
func NewTaobaoQimenStoreitemQueryRequest() *TaobaoQimenStoreitemQueryAPIRequest {
	return &TaobaoQimenStoreitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreitemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Page Setter
// 当前页面
func (r *TaobaoQimenStoreitemQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r TaobaoQimenStoreitemQueryAPIRequest) GetPage() int64 {
	return r._page
}

// Set is StoreId Setter
// 线上门店id
func (r *TaobaoQimenStoreitemQueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoQimenStoreitemQueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
