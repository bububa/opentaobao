package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家仓物流发货推单接单回告 API请求
alibaba.ascp.uop.supplier.consignorder.notify.received

ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
*/
type AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest struct {
    model.Params
    // qimen.alibaba.ascp.uop.consignorder.notify报文中的supplierId字段值
    _supplierId   string
    // qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
    _bizOrderCode   string
    // 业务请求时间
    _bizTime   string
}

// 初始化AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest对象
func NewAlibabaAscpUopSupplierConsignorderNotifyReceivedRequest() *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest{
    return &AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.notify.received"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierId Setter
// qimen.alibaba.ascp.uop.consignorder.notify报文中的supplierId字段值
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetSupplierId(_supplierId string) error {
    r._supplierId = _supplierId
    r.Set("supplier_id", _supplierId)
    return nil
}

// SupplierId Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetSupplierId() string {
    return r._supplierId
}
// BizOrderCode Setter
// qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetBizOrderCode(_bizOrderCode string) error {
    r._bizOrderCode = _bizOrderCode
    r.Set("biz_order_code", _bizOrderCode)
    return nil
}

// BizOrderCode Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetBizOrderCode() string {
    return r._bizOrderCode
}
// BizTime Setter
// 业务请求时间
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetBizTime(_bizTime string) error {
    r._bizTime = _bizTime
    r.Set("biz_time", _bizTime)
    return nil
}

// BizTime Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetBizTime() string {
    return r._bizTime
}
