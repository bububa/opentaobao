package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购服务商发货 API请求
alibaba.idle.isv.order.ship

闲鱼无忧购业务入仓模式下服务商订单发货的接口
*/
type AlibabaIdleIsvOrderShipRequest struct {
    model.Params
    // 订单号
    bizOrderId   string
    // 物流公司
    logisticsCompany   string
    // 运单号
    shipMailNo   string
}

// 初始化AlibabaIdleIsvOrderShipRequest对象
func NewAlibabaIdleIsvOrderShipRequest() *AlibabaIdleIsvOrderShipRequest{
    return &AlibabaIdleIsvOrderShipRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderShipRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.ship"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *AlibabaIdleIsvOrderShipRequest) SetBizOrderId(bizOrderId string) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleIsvOrderShipRequest) GetBizOrderId() string {
    return r.bizOrderId
}
// LogisticsCompany Setter
// 物流公司
func (r *AlibabaIdleIsvOrderShipRequest) SetLogisticsCompany(logisticsCompany string) error {
    r.logisticsCompany = logisticsCompany
    r.Set("logistics_company", logisticsCompany)
    return nil
}

// LogisticsCompany Getter
func (r AlibabaIdleIsvOrderShipRequest) GetLogisticsCompany() string {
    return r.logisticsCompany
}
// ShipMailNo Setter
// 运单号
func (r *AlibabaIdleIsvOrderShipRequest) SetShipMailNo(shipMailNo string) error {
    r.shipMailNo = shipMailNo
    r.Set("ship_mail_no", shipMailNo)
    return nil
}

// ShipMailNo Getter
func (r AlibabaIdleIsvOrderShipRequest) GetShipMailNo() string {
    return r.shipMailNo
}
