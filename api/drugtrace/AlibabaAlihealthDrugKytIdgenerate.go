package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
终端(医疗机构|零售药店)ID生成接口 
alibaba.alihealth.drug.kyt.idgenerate

终端(医疗机构|零售药店)ID生成接口
*/
func AlibabaAlihealthDrugKytIdgenerate(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytIdgenerateRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytIdgenerateAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytIdgenerateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
