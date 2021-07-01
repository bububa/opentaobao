package alihealthpw

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthpw"
)

/* 
申请节点变更回调 
alibaba.alihealth.pw.applynode.update

基金会回调阿里健康更新药品援助申请单的状态
*/
func AlibabaAlihealthPwApplynodeUpdate(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwApplynodeUpdateAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwApplynodeUpdateAPIResponse, error) {
    var resp alihealthpw.AlibabaAlihealthPwApplynodeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
