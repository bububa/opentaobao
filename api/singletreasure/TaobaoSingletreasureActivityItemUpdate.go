package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
更新单品优惠接口 
taobao.singletreasure.activity.item.update

更新单品优惠接口
*/
func TaobaoSingletreasureActivityItemUpdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemUpdateAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityItemUpdateAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityItemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
