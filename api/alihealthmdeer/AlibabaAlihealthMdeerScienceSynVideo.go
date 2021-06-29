package alihealthmdeer

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmdeer"
)

/* 
视频同步【保存/更新】 
alibaba.alihealth.mdeer.science.synVideo

视频同步【保存/更新】
*/
func AlibabaAlihealthMdeerScienceSynVideo(clt *core.SDKClient, req *alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoRequest, session string) (*alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIResponse, error) {
    var resp alihealthmdeer.AlibabaAlihealthMdeerScienceSynVideoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
