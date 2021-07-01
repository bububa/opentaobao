package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询码是否激活 
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活
*/
func AlibabaAlihealthDrugKytQuerycodeactive(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
