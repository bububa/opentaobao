package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
单品搜索人群修改状态 
taobao.simba.serchcrowd.state.batch.update

暂停或启用单品推广搜索人群溢价
*/
func TaobaoSimbaSerchcrowdStateBatchUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdStateBatchUpdateRequest, session string) (*simba.TaobaoSimbaSerchcrowdStateBatchUpdateResponse, error) {
    var resp simba.TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
