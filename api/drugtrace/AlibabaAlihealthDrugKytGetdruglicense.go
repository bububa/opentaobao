package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取药品资质信息 
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。
*/
func AlibabaAlihealthDrugKytGetdruglicense(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetdruglicenseAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetdruglicenseAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytGetdruglicenseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
