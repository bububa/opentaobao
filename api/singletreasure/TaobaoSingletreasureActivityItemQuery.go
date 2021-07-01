package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
查询活动下的优惠信息 
taobao.singletreasure.activity.item.query

分页查询活动下的商品优惠信息
*/
func TaobaoSingletreasureActivityItemQuery(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemQueryAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityItemQueryAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
