package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstoreitemqueryAPIRequest 门店关联商品查询接口 API请求
// taobao.qimen.storeitem.query
//
// 商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
type TaobaoqimenstoreitemqueryAPIRequest struct {
	model.Params
	// 当前页面
	_page int64
	// 线上门店id
	_storeId int64
}

// NewTaobaoqimenstoreitemqueryRequest 初始化TaobaoqimenstoreitemqueryAPIRequest对象
func NewTaobaoqimenstoreitemqueryRequest() *TaobaoqimenstoreitemqueryAPIRequest {
	return &TaobaoqimenstoreitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstoreitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstoreitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstoreitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPage is Page Setter
// 当前页面
func (r *TaobaoqimenstoreitemqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoqimenstoreitemqueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetStoreId is StoreId Setter
// 线上门店id
func (r *TaobaoqimenstoreitemqueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoqimenstoreitemqueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}
