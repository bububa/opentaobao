package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.update APIRequest
taobao.omniorder.store.switchstatus.update

变更门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusUpdateRequest struct {
    model.Params

    // 门店ID
    storeId   int64 

    // 门店发货自提状态
    status   string 

}

func NewTaobaoOmniorderStoreSwitchstatusUpdateRequest() *TaobaoOmniorderStoreSwitchstatusUpdateRequest{
    return &TaobaoOmniorderStoreSwitchstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSwitchstatusUpdateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.switchstatus.update"
}

func (r TaobaoOmniorderStoreSwitchstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSwitchstatusUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreSwitchstatusUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreSwitchstatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoOmniorderStoreSwitchstatusUpdateRequest) GetStatus() string {
    return r.status
}

