package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询家装服务商列表 API请求
taobao.wlb.order.jzpartner.query

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
*/
type TaobaoWlbOrderJzpartnerQueryRequest struct {
    model.Params
    // 淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。
    taobaoTradeId   int64
    // serviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）
    serviceType   int64
}

// 初始化TaobaoWlbOrderJzpartnerQueryRequest对象
func NewTaobaoWlbOrderJzpartnerQueryRequest() *TaobaoWlbOrderJzpartnerQueryRequest{
    return &TaobaoWlbOrderJzpartnerQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzpartnerQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jzpartner.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzpartnerQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaobaoTradeId Setter
// 淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。
func (r *TaobaoWlbOrderJzpartnerQueryRequest) SetTaobaoTradeId(taobaoTradeId int64) error {
    r.taobaoTradeId = taobaoTradeId
    r.Set("taobao_trade_id", taobaoTradeId)
    return nil
}

// TaobaoTradeId Getter
func (r TaobaoWlbOrderJzpartnerQueryRequest) GetTaobaoTradeId() int64 {
    return r.taobaoTradeId
}
// ServiceType Setter
// serviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）
func (r *TaobaoWlbOrderJzpartnerQueryRequest) SetServiceType(serviceType int64) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

// ServiceType Getter
func (r TaobaoWlbOrderJzpartnerQueryRequest) GetServiceType() int64 {
    return r.serviceType
}
