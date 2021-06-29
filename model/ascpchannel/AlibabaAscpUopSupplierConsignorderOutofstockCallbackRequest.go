package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
履约单纬度的仓缺货回告服务 APIRequest
alibaba.ascp.uop.supplier.consignorder.outofstock.callback

商家仓履约单纬度的仓缺货回告接口
*/
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest struct {
    model.Params

    // 缺货回告请求模型
    consignorderOutofstockCallbackRequest   *Consignorderoutofstockcallbackrequest 

}

func NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest() *AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest{
    return &AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.supplier.consignorder.outofstock.callback"
}

func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) SetConsignorderOutofstockCallbackRequest(consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest) error {
    r.consignorderOutofstockCallbackRequest = consignorderOutofstockCallbackRequest
    r.Set("consignorder_outofstock_callback_request", consignorderOutofstockCallbackRequest)
    return nil
}

func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest) GetConsignorderOutofstockCallbackRequest() *Consignorderoutofstockcallbackrequest {
    return r.consignorderOutofstockCallbackRequest
}

