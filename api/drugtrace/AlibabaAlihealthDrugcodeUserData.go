package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
西安杨森同步用户行为接口 
alibaba.alihealth.drugcode.user.data

西安杨森同步用户行为接口
*/
func AlibabaAlihealthDrugcodeUserData(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeUserDataRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeUserDataAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugcodeUserDataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
