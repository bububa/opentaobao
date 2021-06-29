package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
通用查询码接口 
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口
*/
func AlibabaAlihealthDrugCodeCommonListCodeinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}