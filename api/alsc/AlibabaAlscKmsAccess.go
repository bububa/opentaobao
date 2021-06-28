package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
本地生活风控数据接入 
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口
*/
func AlibabaAlscKmsAccess(clt *core.SDKClient, req *alsc.AlibabaAlscKmsAccessRequest, session string) (*alsc.AlibabaAlscKmsAccessAPIResponse, error) {
    var resp alsc.AlibabaAlscKmsAccessAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
