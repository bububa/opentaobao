package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
医院药品子码申请接口 
alibaba.alihealth.drug.code.kyt.yy.applycode

根据父码及所属企业ID生成子码信息
*/
func AlibabaAlihealthDrugCodeKytYyApplycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
