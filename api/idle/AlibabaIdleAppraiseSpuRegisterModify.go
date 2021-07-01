package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
验货宝服务商spu挂载 
alibaba.idle.appraise.spu.register.modify

闲鱼接收回收商spu模板挂载信息
*/
func AlibabaIdleAppraiseSpuRegisterModify(clt *core.SDKClient, req *idle.AlibabaIdleAppraiseSpuRegisterModifyAPIRequest, session string) (*idle.AlibabaIdleAppraiseSpuRegisterModifyAPIResponse, error) {
    var resp idle.AlibabaIdleAppraiseSpuRegisterModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
