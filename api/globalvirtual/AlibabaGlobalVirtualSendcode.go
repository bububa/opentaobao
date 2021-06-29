package globalvirtual

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/globalvirtual"
)

/* 
国际虚拟商品发码服务 
alibaba.global.virtual.sendcode

global virtual send code service
*/
func AlibabaGlobalVirtualSendcode(clt *core.SDKClient, req *globalvirtual.AlibabaGlobalVirtualSendcodeRequest, session string) (*globalvirtual.AlibabaGlobalVirtualSendcodeAPIResponse, error) {
    var resp globalvirtual.AlibabaGlobalVirtualSendcodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
