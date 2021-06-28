package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
取得一个推广组的所有关键词 
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词
*/
func TaobaoSimbaKeywordsbyadgroupidGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsbyadgroupidGetRequest, session string) (*simba.TaobaoSimbaKeywordsbyadgroupidGetAPIResponse, error) {
    var resp simba.TaobaoSimbaKeywordsbyadgroupidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
