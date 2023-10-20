package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryWarehouseQueryAPIRequest 分页查询商家仓信息 API请求
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
type TaobaoInventoryWarehouseQueryAPIRequest struct {
	model.Params
	// 页码
	_pageNo int64
	// 页大小
	_pageSize int64
}

// NewTaobaoInventoryWarehouseQueryRequest 初始化TaobaoInventoryWarehouseQueryAPIRequest对象
func NewTaobaoInventoryWarehouseQueryRequest() *TaobaoInventoryWarehouseQueryAPIRequest {
	return &TaobaoInventoryWarehouseQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryWarehouseQueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryWarehouseQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryWarehouseQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoInventoryWarehouseQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoInventoryWarehouseQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoInventoryWarehouseQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoInventoryWarehouseQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
