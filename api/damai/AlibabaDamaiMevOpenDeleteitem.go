package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票品接口deleteItem 
alibaba.damai.mev.open.deleteitem

deleteItem
*/
func AlibabaDamaiMevOpenDeleteitem(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteitemAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeleteitemAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenDeleteitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
