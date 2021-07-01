package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
修改活动接口 
taobao.singletreasure.activity.update

修改活动接口
*/
func TaobaoSingletreasureActivityUpdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityUpdateAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityUpdateAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
