package usergrowth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth2"
)

/* 
广告投放询问 
taobao.usergrowth.delivery.ask

提供给媒体在曝光广告前调用， 返回是否曝光以及曝光的物料信息
*/
func TaobaoUsergrowthDeliveryAsk(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthDeliveryAskAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthDeliveryAskAPIResponse, error) {
    var resp usergrowth2.TaobaoUsergrowthDeliveryAskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
