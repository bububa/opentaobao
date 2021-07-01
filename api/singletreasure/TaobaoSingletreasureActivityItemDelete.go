package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
删除单品优惠接口 
taobao.singletreasure.activity.item.delete

删除单品优惠接口
*/
func TaobaoSingletreasureActivityItemDelete(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemDeleteAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityItemDeleteAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityItemDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
