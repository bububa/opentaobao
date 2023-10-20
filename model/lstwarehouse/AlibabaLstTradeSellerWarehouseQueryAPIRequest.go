package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerwarehousequeryAPIRequest 供应商-本云商家-仓库查询接口 API请求
// alibaba.lst.trade.seller.warehouse.query
//
// 查询本地云仓商家的仓库
type AlibabalsttradesellerwarehousequeryAPIRequest struct {
	model.Params
	// 入参
	_warehouseQueryParam *WarehouseQueryParam
}

// NewAlibabalsttradesellerwarehousequeryRequest 初始化AlibabalsttradesellerwarehousequeryAPIRequest对象
func NewAlibabalsttradesellerwarehousequeryRequest() *AlibabalsttradesellerwarehousequeryAPIRequest {
	return &AlibabalsttradesellerwarehousequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerwarehousequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerwarehousequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerwarehousequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseQueryParam is WarehouseQueryParam Setter
// 入参
func (r *AlibabalsttradesellerwarehousequeryAPIRequest) SetWarehouseQueryParam(_warehouseQueryParam *WarehouseQueryParam) error {
	r._warehouseQueryParam = _warehouseQueryParam
	r.Set("warehouse_query_param", _warehouseQueryParam)
	return nil
}

// GetWarehouseQueryParam WarehouseQueryParam Getter
func (r AlibabalsttradesellerwarehousequeryAPIRequest) GetWarehouseQueryParam() *WarehouseQueryParam {
	return r._warehouseQueryParam
}
