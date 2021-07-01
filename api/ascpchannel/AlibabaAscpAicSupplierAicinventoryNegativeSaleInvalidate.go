package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
负卖库存失效接口 
alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate

失效负卖库存数据
*/
func AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest, session string) (*ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
