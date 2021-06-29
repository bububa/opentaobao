package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询药监码对应的有效期和包装规格 
alibaba.alihealth.drug.scan.querycode

查询药监码对应的有效期和包装规格
*/
func AlibabaAlihealthDrugScanQuerycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugScanQuerycodeRequest, session string) (*drugtrace.AlibabaAlihealthDrugScanQuerycodeAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugScanQuerycodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
