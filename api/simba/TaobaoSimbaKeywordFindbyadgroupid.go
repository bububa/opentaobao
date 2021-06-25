package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
根据推广单元id获取关键词 
taobao.simba.keyword.findbyadgroupid

根据一个关键词Id列表取得一组关键词
*/
func TaobaoSimbaKeywordFindbyadgroupid(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordFindbyadgroupidRequest, session string) (*simba.TaobaoSimbaKeywordFindbyadgroupidResponse, error) {
    var resp simba.TaobaoSimbaKeywordFindbyadgroupidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
