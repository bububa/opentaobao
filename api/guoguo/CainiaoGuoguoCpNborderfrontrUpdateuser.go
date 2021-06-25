package guoguo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/guoguo"
)

/* 
小件员信息变更 
cainiao.guoguo.cp.nborderfrontr.updateuser

小件员信息变更
*/
func CainiaoGuoguoCpNborderfrontrUpdateuser(clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpNborderfrontrUpdateuserRequest, session string) (*guoguo.CainiaoGuoguoCpNborderfrontrUpdateuserResponse, error) {
    var resp guoguo.CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
