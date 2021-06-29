package sungari

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/sungari"
)

/* 
商品商家处置结果查询 
taobao.sungari.dispose.query

红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
*/
func TaobaoSungariDisposeQuery(clt *core.SDKClient, req *sungari.TaobaoSungariDisposeQueryRequest, session string) (*sungari.TaobaoSungariDisposeQueryAPIResponse, error) {
    var resp sungari.TaobaoSungariDisposeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
