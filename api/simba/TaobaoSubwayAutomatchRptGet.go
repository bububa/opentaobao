package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
查询流量智选天级报告 
taobao.subway.automatch.rpt.get

查询流量智选天级报告
*/
func TaobaoSubwayAutomatchRptGet(clt *core.SDKClient, req *simba.TaobaoSubwayAutomatchRptGetRequest, session string) (*simba.TaobaoSubwayAutomatchRptGetResponse, error) {
    var resp simba.TaobaoSubwayAutomatchRptGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
