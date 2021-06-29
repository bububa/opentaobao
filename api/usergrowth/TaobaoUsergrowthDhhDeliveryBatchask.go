package usergrowth

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth"
)

/* 
广告曝光前判定批量接口V2 
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2
*/
func TaobaoUsergrowthDhhDeliveryBatchask(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDhhDeliveryBatchaskRequest, session string) (*usergrowth.TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse, error) {
    var resp usergrowth.TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
