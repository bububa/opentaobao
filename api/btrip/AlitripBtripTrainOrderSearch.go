package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
获取火车票订单列表 
alitrip.btrip.train.order.search

第三方OA厂商获取自己的火车票数据
*/
func AlitripBtripTrainOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripTrainOrderSearchRequest, session string) (*btrip.AlitripBtripTrainOrderSearchAPIResponse, error) {
    var resp btrip.AlitripBtripTrainOrderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
