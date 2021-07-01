package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
商家推单 
alibaba.ascp.industry.uop.supplier.consignoder

商家推单
*/
func AlibabaAscpIndustryUopSupplierConsignoder(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIRequest, session string) (*ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
