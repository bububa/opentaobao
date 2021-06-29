package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售商家仓出库 API请求
alibaba.ascp.uop.taobao.presalesorder.consignconfirm

预售商家仓出库
*/
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest struct {
    model.Params
    // 预售订单商家仓出库对象
    presalesOrderConsignConfirmRequest   *Presalesorderconsignconfirmrequest
}

// 初始化AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest对象
func NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest() *AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest{
    return &AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.taobao.presalesorder.consignconfirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PresalesOrderConsignConfirmRequest Setter
// 预售订单商家仓出库对象
func (r *AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest) SetPresalesOrderConsignConfirmRequest(presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest) error {
    r.presalesOrderConsignConfirmRequest = presalesOrderConsignConfirmRequest
    r.Set("presales_order_consign_confirm_request", presalesOrderConsignConfirmRequest)
    return nil
}

// PresalesOrderConsignConfirmRequest Getter
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest) GetPresalesOrderConsignConfirmRequest() *Presalesorderconsignconfirmrequest {
    return r.presalesOrderConsignConfirmRequest
}
