package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取盲底文件处理结果 
alibaba.alihealth.drugcode.drugfactory.getblindresult

获取盲底文件处理结果
*/
func AlibabaAlihealthDrugcodeDrugfactoryGetblindresult(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
