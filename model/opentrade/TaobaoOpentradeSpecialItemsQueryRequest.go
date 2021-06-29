package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单获取商品绑定信息 APIRequest
taobao.opentrade.special.items.query

专属下单获取商品绑定信息
*/
type TaobaoOpentradeSpecialItemsQueryRequest struct {
    model.Params

    // 绑定专属下单场景的C端小程序ID
    miniappId   int64 

}

func NewTaobaoOpentradeSpecialItemsQueryRequest() *TaobaoOpentradeSpecialItemsQueryRequest{
    return &TaobaoOpentradeSpecialItemsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeSpecialItemsQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.query"
}

func (r TaobaoOpentradeSpecialItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeSpecialItemsQueryRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeSpecialItemsQueryRequest) GetMiniappId() int64 {
    return r.miniappId
}

