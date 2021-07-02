package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzpartnerQueryAPIRequest 查询家装服务商列表 API请求
// taobao.wlb.order.jzpartner.query
//
// 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryAPIRequest struct {
	model.Params
	// 淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。
	_taobaoTradeId int64
	// serviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）
	_serviceType int64
}

// NewTaobaoWlbOrderJzpartnerQueryRequest 初始化TaobaoWlbOrderJzpartnerQueryAPIRequest对象
func NewTaobaoWlbOrderJzpartnerQueryRequest() *TaobaoWlbOrderJzpartnerQueryAPIRequest {
	return &TaobaoWlbOrderJzpartnerQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzpartnerQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jzpartner.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzpartnerQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTaobaoTradeId is TaobaoTradeId Setter
// 淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。
func (r *TaobaoWlbOrderJzpartnerQueryAPIRequest) SetTaobaoTradeId(_taobaoTradeId int64) error {
	r._taobaoTradeId = _taobaoTradeId
	r.Set("taobao_trade_id", _taobaoTradeId)
	return nil
}

// GetTaobaoTradeId TaobaoTradeId Getter
func (r TaobaoWlbOrderJzpartnerQueryAPIRequest) GetTaobaoTradeId() int64 {
	return r._taobaoTradeId
}

// SetServiceType is ServiceType Setter
// serviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）
func (r *TaobaoWlbOrderJzpartnerQueryAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TaobaoWlbOrderJzpartnerQueryAPIRequest) GetServiceType() int64 {
	return r._serviceType
}
