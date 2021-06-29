package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单商家仓发货结果回传服务 API请求
alibaba.ascp.uop.supplier.consignorder.ship

ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
*/
type AlibabaAscpUopSupplierConsignorderShipRequest struct {
    model.Params
    // 发货回传请求模型
    consignorderShipRequest   *Consignordershiprequest
}

// 初始化AlibabaAscpUopSupplierConsignorderShipRequest对象
func NewAlibabaAscpUopSupplierConsignorderShipRequest() *AlibabaAscpUopSupplierConsignorderShipRequest{
    return &AlibabaAscpUopSupplierConsignorderShipRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.ship"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignorderShipRequest Setter
// 发货回传请求模型
func (r *AlibabaAscpUopSupplierConsignorderShipRequest) SetConsignorderShipRequest(consignorderShipRequest *Consignordershiprequest) error {
    r.consignorderShipRequest = consignorderShipRequest
    r.Set("consignorder_ship_request", consignorderShipRequest)
    return nil
}

// ConsignorderShipRequest Getter
func (r AlibabaAscpUopSupplierConsignorderShipRequest) GetConsignorderShipRequest() *Consignordershiprequest {
    return r.consignorderShipRequest
}
