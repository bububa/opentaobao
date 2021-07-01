package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦电商对外搜索服务 
alibaba.damai.ec.search.project.search

大麦电商对外搜索服务
*/
func AlibabaDamaiEcSearchProjectSearch(clt *core.SDKClient, req *damai.AlibabaDamaiEcSearchProjectSearchAPIRequest, session string) (*damai.AlibabaDamaiEcSearchProjectSearchAPIResponse, error) {
    var resp damai.AlibabaDamaiEcSearchProjectSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
