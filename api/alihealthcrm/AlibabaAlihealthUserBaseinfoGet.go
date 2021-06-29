package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
获取用户基础信息 
alibaba.alihealth.user.baseinfo.get

获取用户基础信息
*/
func AlibabaAlihealthUserBaseinfoGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthUserBaseinfoGetRequest, session string) (*alihealthcrm.AlibabaAlihealthUserBaseinfoGetAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthUserBaseinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}