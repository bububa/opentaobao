package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
批量修改商品接口 
taobao.singletreasure.activity.item.batchupdate

批量修改商品优惠接口
*/
func TaobaoSingletreasureActivityItemBatchupdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemBatchupdateAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityItemBatchupdateAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityItemBatchupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
