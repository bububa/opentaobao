package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
获取产品下挂载的技能列表 
taobao.ailab.aicloud.top.skils.list.new

星空平台提供的获取产品下挂载的技能列表新接口
*/
func TaobaoAilabAicloudTopSkilsListNew(clt *core.SDKClient, req *alilabs.TaobaoAilabAicloudTopSkilsListNewAPIRequest, session string) (*alilabs.TaobaoAilabAicloudTopSkilsListNewAPIResponse, error) {
    var resp alilabs.TaobaoAilabAicloudTopSkilsListNewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
