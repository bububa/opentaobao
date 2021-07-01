package usergrowth

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth"
)

/* 
广告投放批量询问 
taobao.usergrowth.delivery.batchask

提供给媒体在曝光广告前调用， 返回是否曝光以及报价
*/
func TaobaoUsergrowthDeliveryBatchask(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDeliveryBatchaskAPIRequest, session string) (*usergrowth.TaobaoUsergrowthDeliveryBatchaskAPIResponse, error) {
    var resp usergrowth.TaobaoUsergrowthDeliveryBatchaskAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
