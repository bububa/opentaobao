package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购服务商发货 APIRequest
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

func NewAlibabaIdleIsvOrderShipRequest() *AlibabaIdleIsvOrderShipRequest{
    return &AlibabaIdleIsvOrderShipRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvOrderShipRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.ship"
}

func (r AlibabaIdleIsvOrderShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvOrderShipRequest) SetBizOrderId(bizOrderId string) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaIdleIsvOrderShipRequest) GetBizOrderId() string {
    return r.bizOrderId
}

func (r *AlibabaIdleIsvOrderShipRequest) SetLogisticsCompany(logisticsCompany string) error {
    r.logisticsCompany = logisticsCompany
    r.Set("logistics_company", logisticsCompany)
    return nil
}

func (r AlibabaIdleIsvOrderShipRequest) GetLogisticsCompany() string {
    return r.logisticsCompany
}

func (r *AlibabaIdleIsvOrderShipRequest) SetShipMailNo(shipMailNo string) error {
    r.shipMailNo = shipMailNo
    r.Set("ship_mail_no", shipMailNo)
    return nil
}

func (r AlibabaIdleIsvOrderShipRequest) GetShipMailNo() string {
    return r.shipMailNo
}

