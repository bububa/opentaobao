package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
活动创建接口 
taobao.singletreasure.activity.create

创建优惠活动
*/
func TaobaoSingletreasureActivityCreate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityCreateRequest, session string) (*singletreasure.TaobaoSingletreasureActivityCreateAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
