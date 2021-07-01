package usergrowth

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth"
)

/* 
广告曝光前判定接口V2 
taobao.usergrowth.dhh.delivery.ask

提供给媒体在曝光广告前调用
*/
func TaobaoUsergrowthDhhDeliveryAsk(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIRequest, session string) (*usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIResponse, error) {
    var resp usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
