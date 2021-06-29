package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
零售端平台单据查询 
alibaba.alihealth.drug.kyt.storebilllist

零售端平台单据查询
*/
func AlibabaAlihealthDrugKytStorebilllist(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytStorebilllistRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytStorebilllistAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytStorebilllistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
