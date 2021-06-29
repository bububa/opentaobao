package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取企业资质 
alibaba.alihealth.drug.kyt.getentlicense

获取企业的资质信息。
*/
func AlibabaAlihealthDrugKytGetentlicense(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetentlicenseRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
