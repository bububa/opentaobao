package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询负卖库存 APIRequest
alibaba.ascp.aic.supplier.aicinventory.negative.sale.query

商家根据当前接口查询负卖货品的库存
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest struct {
    model.Params

    // 库存查询参数
    aicinventoryQueryRequest   *Aicinventoryqueryrequest 

}

func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.query"
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) SetAicinventoryQueryRequest(aicinventoryQueryRequest *Aicinventoryqueryrequest) error {
    r.aicinventoryQueryRequest = aicinventoryQueryRequest
    r.Set("aicinventory_query_request", aicinventoryQueryRequest)
    return nil
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetAicinventoryQueryRequest() *Aicinventoryqueryrequest {
    return r.aicinventoryQueryRequest
}

