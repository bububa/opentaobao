package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
获取神鲸用户基本信息 
alibaba.iwork.core.hrs.getperson

神鲸用户的基本信息查询，根据PERSON_ID或者用户ACCOUNT_ID查询
*/
func AlibabaIworkCoreHrsGetperson(clt *core.SDKClient, req *campus.AlibabaIworkCoreHrsGetpersonAPIRequest, session string) (*campus.AlibabaIworkCoreHrsGetpersonAPIResponse, error) {
    var resp campus.AlibabaIworkCoreHrsGetpersonAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
