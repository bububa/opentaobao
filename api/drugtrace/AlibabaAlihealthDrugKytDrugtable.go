package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询药品目录信息 
alibaba.alihealth.drug.kyt.drugtable

查询药品目录信息
*/
func AlibabaAlihealthDrugKytDrugtable(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugtableRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrugtableAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDrugtableAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
