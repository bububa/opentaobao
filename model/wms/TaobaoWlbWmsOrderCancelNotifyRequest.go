package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 API请求
taobao.wlb.wms.order.cancel.notify

单据取消接口
*/
type TaobaoWlbWmsOrderCancelNotifyRequest struct {
    model.Params
    // 订单类型
    orderCode   string
    // 仓库编码
    storeCode   string
    // 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
    orderType   string
}

// 初始化TaobaoWlbWmsOrderCancelNotifyRequest对象
func NewTaobaoWlbWmsOrderCancelNotifyRequest() *TaobaoWlbWmsOrderCancelNotifyRequest{
    return &TaobaoWlbWmsOrderCancelNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsOrderCancelNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.order.cancel.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsOrderCancelNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 订单类型
func (r *TaobaoWlbWmsOrderCancelNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsOrderCancelNotifyRequest) GetOrderCode() string {
    return r.orderCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsOrderCancelNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsOrderCancelNotifyRequest) GetStoreCode() string {
    return r.storeCode
}
// OrderType Setter
// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
func (r *TaobaoWlbWmsOrderCancelNotifyRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsOrderCancelNotifyRequest) GetOrderType() string {
    return r.orderType
}
