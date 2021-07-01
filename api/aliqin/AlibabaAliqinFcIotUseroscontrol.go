package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
物联卡用户管理停开机功能 
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持
*/
func AlibabaAliqinFcIotUseroscontrol(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotUseroscontrolAPIRequest, session string) (*aliqin.AlibabaAliqinFcIotUseroscontrolAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotUseroscontrolAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
