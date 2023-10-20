package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbitembatchqueryAPIRequest 批次库存查询接口 API请求
// taobao.wlb.item.batch.query
//
// 根据用户id，item id list和store code来查询商品库存信息和批次信息
type TaobaowlbitembatchqueryAPIRequest struct {
	model.Params
	// 需要查询的商品ID列表，以字符串表示，ID间以;隔开
	_itemIds string
	// 仓库编号
	_storeCode string
	// 分页查询参数，指定查询页数，默认为1
	_pageNo int64
	// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
	_pageSize int64
}

// NewTaobaowlbitembatchqueryRequest 初始化TaobaowlbitembatchqueryAPIRequest对象
func NewTaobaowlbitembatchqueryRequest() *TaobaowlbitembatchqueryAPIRequest {
	return &TaobaowlbitembatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbitembatchqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.item.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbitembatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbitembatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 需要查询的商品ID列表，以字符串表示，ID间以;隔开
func (r *TaobaowlbitembatchqueryAPIRequest) SetItemIds(_itemIds string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaowlbitembatchqueryAPIRequest) GetItemIds() string {
	return r._itemIds
}

// SetStoreCode is StoreCode Setter
// 仓库编号
func (r *TaobaowlbitembatchqueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbitembatchqueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetPageNo is PageNo Setter
// 分页查询参数，指定查询页数，默认为1
func (r *TaobaowlbitembatchqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaowlbitembatchqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
func (r *TaobaowlbitembatchqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaowlbitembatchqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
