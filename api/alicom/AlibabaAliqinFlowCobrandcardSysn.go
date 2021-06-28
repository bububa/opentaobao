package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
联名卡信息同步 
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。
*/
func AlibabaAliqinFlowCobrandcardSysn(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowCobrandcardSysnRequest, session string) (*alicom.AlibabaAliqinFlowCobrandcardSysnAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowCobrandcardSysnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
