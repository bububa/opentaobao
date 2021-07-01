package jstinteractive

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jstinteractive"
)

/* 
查询可配置任务素材接口 
taobao.jst.interactive.assets.query

查询可配置任务素材列表，用以配置任务素材
*/
func TaobaoJstInteractiveAssetsQuery(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsQueryAPIRequest, session string) (*jstinteractive.TaobaoJstInteractiveAssetsQueryAPIResponse, error) {
    var resp jstinteractive.TaobaoJstInteractiveAssetsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
