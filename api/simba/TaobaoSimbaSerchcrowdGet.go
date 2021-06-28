package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
根据推广单元id获取搜索溢价人群 
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群
*/
func TaobaoSimbaSerchcrowdGet(clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdGetRequest, session string) (*simba.TaobaoSimbaSerchcrowdGetAPIResponse, error) {
    var resp simba.TaobaoSimbaSerchcrowdGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
