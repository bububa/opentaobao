package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
购物金卡查询 
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。
*/
func TaobaoCardExpandcardQuery(clt *core.SDKClient, req *promotion.TaobaoCardExpandcardQueryRequest, session string) (*promotion.TaobaoCardExpandcardQueryAPIResponse, error) {
    var resp promotion.TaobaoCardExpandcardQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
