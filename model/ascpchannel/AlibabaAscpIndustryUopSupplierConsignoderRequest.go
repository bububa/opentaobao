package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家推单 APIRequest
alibaba.ascp.industry.uop.supplier.consignoder

商家推单
*/
type AlibabaAscpIndustryUopSupplierConsignoderRequest struct {
    model.Params

    // 发货主单信息
    erpNormalConsignOrderRequest   *Erpnormalconsignorderrequest 

}

func NewAlibabaAscpIndustryUopSupplierConsignoderRequest() *AlibabaAscpIndustryUopSupplierConsignoderRequest{
    return &AlibabaAscpIndustryUopSupplierConsignoderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.uop.supplier.consignoder"
}

func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpIndustryUopSupplierConsignoderRequest) SetErpNormalConsignOrderRequest(erpNormalConsignOrderRequest *Erpnormalconsignorderrequest) error {
    r.erpNormalConsignOrderRequest = erpNormalConsignOrderRequest
    r.Set("erp_normal_consign_order_request", erpNormalConsignOrderRequest)
    return nil
}

func (r AlibabaAscpIndustryUopSupplierConsignoderRequest) GetErpNormalConsignOrderRequest() *Erpnormalconsignorderrequest {
    return r.erpNormalConsignOrderRequest
}

