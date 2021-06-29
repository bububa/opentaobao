package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询负卖库存 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.query

商家根据当前接口查询负卖货品的库存
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest struct {
    model.Params
    // 库存查询参数
    aicinventoryQueryRequest   *Aicinventoryqueryrequest
}

// 初始化AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AicinventoryQueryRequest Setter
// 库存查询参数
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) SetAicinventoryQueryRequest(aicinventoryQueryRequest *Aicinventoryqueryrequest) error {
    r.aicinventoryQueryRequest = aicinventoryQueryRequest
    r.Set("aicinventory_query_request", aicinventoryQueryRequest)
    return nil
}

// AicinventoryQueryRequest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryRequest) GetAicinventoryQueryRequest() *Aicinventoryqueryrequest {
    return r.aicinventoryQueryRequest
}
