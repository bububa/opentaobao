package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuwarehouseskuscrollqueryAPIRequest 仓商品遍历接口(游标) API请求
// alibaba.wdk.sku.warehousesku.scroll.query
//
// 提供仓商品数据接口查询
type AlibabawdkskuwarehouseskuscrollqueryAPIRequest struct {
	model.Params
	// 仓库编码
	_warehouseCode string
	// 游标
	_scrollId string
}

// NewAlibabawdkskuwarehouseskuscrollqueryRequest 初始化AlibabawdkskuwarehouseskuscrollqueryAPIRequest对象
func NewAlibabawdkskuwarehouseskuscrollqueryRequest() *AlibabawdkskuwarehouseskuscrollqueryAPIRequest {
	return &AlibabawdkskuwarehouseskuscrollqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuwarehouseskuscrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.warehousesku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuwarehouseskuscrollqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuwarehouseskuscrollqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库编码
func (r *AlibabawdkskuwarehouseskuscrollqueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkskuwarehouseskuscrollqueryAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetScrollId is ScrollId Setter
// 游标
func (r *AlibabawdkskuwarehouseskuscrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabawdkskuwarehouseskuscrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
