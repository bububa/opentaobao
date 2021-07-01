package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
客户账户报表基础数据对象 
taobao.simba.rpt.custbase.get

客户账户报表基础数据对象
*/
func TaobaoSimbaRptCustbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCustbaseGetAPIRequest, session string) (*simba.TaobaoSimbaRptCustbaseGetAPIResponse, error) {
    var resp simba.TaobaoSimbaRptCustbaseGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
