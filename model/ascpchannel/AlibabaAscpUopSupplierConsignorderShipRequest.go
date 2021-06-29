package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单商家仓发货结果回传服务 APIRequest
alibaba.ascp.uop.supplier.consignorder.ship

ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
*/
type AlibabaAscpUopSupplierConsignorderShipRequest struct {
    model.Params

    // 发货回传请求模型
    consignorderShipRequest   *Consignordershiprequest 

}

func NewAlibabaAscpUopSupplierConsignorderShipRequest() *AlibabaAscpUopSupplierConsignorderShipRequest{
    return &AlibabaAscpUopSupplierConsignorderShipRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.ship"
}

func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierConsignorderShipRequest) SetConsignorderShipRequest(consignorderShipRequest *Consignordershiprequest) error {
    r.consignorderShipRequest = consignorderShipRequest
    r.Set("consignorder_ship_request", consignorderShipRequest)
    return nil
}

func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetConsignorderShipRequest() *Consignordershiprequest {
    return r.consignorderShipRequest
}

