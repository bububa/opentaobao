package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
查询物联卡个人实人认证信息 
alibaba.aliqin.fc.iot.qry.personinfo

查询物联卡个人实人认证信息
*/
func AlibabaAliqinFcIotQryPersoninfo(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQryPersoninfoAPIRequest, session string) (*aliqin.AlibabaAliqinFcIotQryPersoninfoAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotQryPersoninfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
