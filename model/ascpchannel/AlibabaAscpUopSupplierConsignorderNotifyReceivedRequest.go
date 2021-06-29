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
    supplierId   string
    // qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
    bizOrderCode   string
    // 业务请求时间
    bizTime   string
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
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetSupplierId(supplierId string) error {
    r.supplierId = supplierId
    r.Set("supplier_id", supplierId)
    return nil
}

// SupplierId Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetSupplierId() string {
    return r.supplierId
}
// BizOrderCode Setter
// qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetBizOrderCode(bizOrderCode string) error {
    r.bizOrderCode = bizOrderCode
    r.Set("biz_order_code", bizOrderCode)
    return nil
}

// BizOrderCode Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetBizOrderCode() string {
    return r.bizOrderCode
}
// BizTime Setter
// 业务请求时间
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) SetBizTime(bizTime string) error {
    r.bizTime = bizTime
    r.Set("biz_time", bizTime)
    return nil
}

// BizTime Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedRequest) GetBizTime() string {
    return r.bizTime
}
