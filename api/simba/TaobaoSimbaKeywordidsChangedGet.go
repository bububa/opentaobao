package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取修改的词ID 
taobao.simba.keywordids.changed.get

获取修改的词ID
*/
func TaobaoSimbaKeywordidsChangedGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordidsChangedGetAPIRequest, session string) (*simba.TaobaoSimbaKeywordidsChangedGetAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordidsChangedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
