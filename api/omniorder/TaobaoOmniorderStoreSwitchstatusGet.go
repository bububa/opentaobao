package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
switchstatus.get 
taobao.omniorder.store.switchstatus.get

查询门店发货、门店自提状态
*/
func TaobaoOmniorderStoreSwitchstatusGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSwitchstatusGetRequest, session string) (*omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreSwitchstatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
