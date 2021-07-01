package alihealthmdeer

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmdeer"
)

/* 
医知鹿文章同步【保存/更新】 
alibaba.alihealth.mdeer.science.synarticle

文章同步【保存/更新】
*/
func AlibabaAlihealthMdeerScienceSynarticle(clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceSynarticleAPIRequest, session string) (*alihealthmdeer.AlibabaAlihealthMdeerScienceSynarticleAPIResponse, error) {
    var resp alihealthmdeer.AlibabaAlihealthMdeerScienceSynarticleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
