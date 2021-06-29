package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放获取商品绑定信息 APIRequest
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息
*/
type TaobaoOpentradeToolsItemsQueryRequest struct {
    model.Params

    // 交易开放C端小程序ID
    miniappId   int64 

}

func NewTaobaoOpentradeToolsItemsQueryRequest() *TaobaoOpentradeToolsItemsQueryRequest{
    return &TaobaoOpentradeToolsItemsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeToolsItemsQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.tools.items.query"
}

func (r TaobaoOpentradeToolsItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeToolsItemsQueryRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeToolsItemsQueryRequest) GetMiniappId() int64 {
    return r.miniappId
}

