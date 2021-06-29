package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化接口 APIRequest
taobao.jst.astrolabe.storeinventory.iteminitial

ERP调用奇门的接口，对门店的库存进行初始化
*/
type TaobaoJstAstrolabeStoreinventoryIteminitialRequest struct {
    model.Params

    // 门店列表
    stores   []Store 

    // 操作时间
    operationTime   string 

}

func NewTaobaoJstAstrolabeStoreinventoryIteminitialRequest() *TaobaoJstAstrolabeStoreinventoryIteminitialRequest{
    return &TaobaoJstAstrolabeStoreinventoryIteminitialRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.iteminitial"
}

func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstAstrolabeStoreinventoryIteminitialRequest) SetStores(stores []Store) error {
    r.stores = stores
    r.Set("stores", stores)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetStores() []Store {
    return r.stores
}

func (r *TaobaoJstAstrolabeStoreinventoryIteminitialRequest) SetOperationTime(operationTime string) error {
    r.operationTime = operationTime
    r.Set("operation_time", operationTime)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryIteminitialRequest) GetOperationTime() string {
    return r.operationTime
}

