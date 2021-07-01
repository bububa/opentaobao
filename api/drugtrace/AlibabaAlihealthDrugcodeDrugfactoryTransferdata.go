package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
药厂传输数据 
alibaba.alihealth.drugcode.drugfactory.transferdata

药厂传输数据
*/
func AlibabaAlihealthDrugcodeDrugfactoryTransferdata(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
