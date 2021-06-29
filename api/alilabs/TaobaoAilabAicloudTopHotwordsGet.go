package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
获取热词 
taobao.ailab.aicloud.top.hotwords.get

获取ASR热词
*/
func TaobaoAilabAicloudTopHotwordsGet(clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopHotwordsGetRequest, session string) (*alilabs.TaobaoAilabAicloudTopHotwordsGetAPIResponse, error) {
    var resp alilabs.TaobaoAilabAicloudTopHotwordsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
