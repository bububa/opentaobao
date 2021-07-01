package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
更新热词 
taobao.ailab.aicloud.top.hotwords.update

更新ASR热词
*/
func TaobaoAilabAicloudTopHotwordsUpdate(clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIRequest, session string) (*alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIResponse, error) {
    var resp alilabs.TaobaoAilabAicloudTopHotwordsUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
