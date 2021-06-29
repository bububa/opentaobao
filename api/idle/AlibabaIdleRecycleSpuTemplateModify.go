package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
闲鱼接收回收商spu模板挂载信息 
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息
*/
func AlibabaIdleRecycleSpuTemplateModify(clt *core.SDKClient, req *idle.AlibabaIdleRecycleSpuTemplateModifyRequest, session string) (*idle.AlibabaIdleRecycleSpuTemplateModifyAPIResponse, error) {
    var resp idle.AlibabaIdleRecycleSpuTemplateModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
