package alihealthmdeer

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmdeer"
)

/* 
文章删除 
alibaba.alihealth.mdeer.science.deletearticle

三方同步文章删除
*/
func AlibabaAlihealthMdeerScienceDeletearticle(clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceDeletearticleAPIRequest, session string) (*alihealthmdeer.AlibabaAlihealthMdeerScienceDeletearticleAPIResponse, error) {
    var resp alihealthmdeer.AlibabaAlihealthMdeerScienceDeletearticleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
