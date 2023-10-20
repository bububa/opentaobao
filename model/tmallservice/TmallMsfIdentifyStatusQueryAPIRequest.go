package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmsfidentifystatusqueryAPIRequest 喵师傅定案核销状态查询 API请求
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
type TmallmsfidentifystatusqueryAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
	_serviceType int64
}

// NewTmallmsfidentifystatusqueryRequest 初始化TmallmsfidentifystatusqueryAPIRequest对象
func NewTmallmsfidentifystatusqueryRequest() *TmallmsfidentifystatusqueryAPIRequest {
	return &TmallmsfidentifystatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmsfidentifystatusqueryAPIRequest) GetApiMethodName() string {
	return "tmall.msf.identify.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmsfidentifystatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmsfidentifystatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 天猫订单号
func (r *TmallmsfidentifystatusqueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallmsfidentifystatusqueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetServiceType is ServiceType Setter
// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
func (r *TmallmsfidentifystatusqueryAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TmallmsfidentifystatusqueryAPIRequest) GetServiceType() int64 {
	return r._serviceType
}
