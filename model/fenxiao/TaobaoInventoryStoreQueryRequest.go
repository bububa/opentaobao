package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库信息 APIRequest
taobao.inventory.store.query

查询商家仓信息
*/
type TaobaoInventoryStoreQueryRequest struct {
    model.Params

    // 商家的仓库编码
    storeCode   string 

}

func NewTaobaoInventoryStoreQueryRequest() *TaobaoInventoryStoreQueryRequest{
    return &TaobaoInventoryStoreQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryStoreQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.store.query"
}

func (r TaobaoInventoryStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryStoreQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoInventoryStoreQueryRequest) GetStoreCode() string {
    return r.storeCode
}

