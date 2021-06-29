package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.get APIRequest
taobao.omniorder.store.switchstatus.get

查询门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusGetRequest struct {
    model.Params

    // 门店ID
    storeId   int64 

    // 卖家ID
    sellerId   int64 

}

func NewTaobaoOmniorderStoreSwitchstatusGetRequest() *TaobaoOmniorderStoreSwitchstatusGetRequest{
    return &TaobaoOmniorderStoreSwitchstatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.switchstatus.get"
}

func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSwitchstatusGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreSwitchstatusGetRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoOmniorderStoreSwitchstatusGetRequest) GetSellerId() int64 {
    return r.sellerId
}

