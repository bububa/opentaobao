package singletreasure

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/singletreasure"
)

/* 
查询活动列表接口 
taobao.singletreasure.activity.query

查询活动列表接口
*/
func TaobaoSingletreasureActivityQuery(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityQueryRequest, session string) (*singletreasure.TaobaoSingletreasureActivityQueryAPIResponse, error) {
    var resp singletreasure.TaobaoSingletreasureActivityQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
