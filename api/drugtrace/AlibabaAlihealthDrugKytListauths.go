package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
企业搜索自己授权的物流企业 
alibaba.alihealth.drug.kyt.listauths

企业搜索自己授权的物流企业
*/
func AlibabaAlihealthDrugKytListauths(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListauthsAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytListauthsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytListauthsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
