package caipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/caipiao"
)

/* 
根据卖家id与appkey获取商品信息 
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。
*/
func TaobaoCaipiaoGoodsInfoGet(clt *core.SDKClient, req *caipiao.TaobaoCaipiaoGoodsInfoGetAPIRequest, session string) (*caipiao.TaobaoCaipiaoGoodsInfoGetAPIResponse, error) {
    var resp caipiao.TaobaoCaipiaoGoodsInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
