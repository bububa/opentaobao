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
    _bizOrderId   string
    // 物流公司
    _logisticsCompany   string
    // 运单号
    _shipMailNo   string
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
func (r *AlibabaIdleIsvOrderShipRequest) SetBizOrderId(_bizOrderId string) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaIdleIsvOrderShipRequest) GetBizOrderId() string {
    return r._bizOrderId
}
// LogisticsCompany Setter
// 物流公司
func (r *AlibabaIdleIsvOrderShipRequest) SetLogisticsCompany(_logisticsCompany string) error {
    r._logisticsCompany = _logisticsCompany
    r.Set("logistics_company", _logisticsCompany)
    return nil
}

// LogisticsCompany Getter
func (r AlibabaIdleIsvOrderShipRequest) GetLogisticsCompany() string {
    return r._logisticsCompany
}
// ShipMailNo Setter
// 运单号
func (r *AlibabaIdleIsvOrderShipRequest) SetShipMailNo(_shipMailNo string) error {
    r._shipMailNo = _shipMailNo
    r.Set("ship_mail_no", _shipMailNo)
    return nil
}

// ShipMailNo Getter
func (r AlibabaIdleIsvOrderShipRequest) GetShipMailNo() string {
    return r._shipMailNo
}
