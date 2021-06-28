package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-权益物料精选 
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
*/
func TaobaoTbkDgOptimusPromotion(clt *core.SDKClient, req *tbk.TaobaoTbkDgOptimusPromotionRequest, session string) (*tbk.TaobaoTbkDgOptimusPromotionAPIResponse, error) {
    var resp tbk.TaobaoTbkDgOptimusPromotionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
