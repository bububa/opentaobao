package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品仓库库存 APIRequest
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
type TmallInventoryQueryForstoreRequest struct {
    model.Params

    // 查询列表
    paramList   []InventoryQueryForStoreRequest 

}

func NewTmallInventoryQueryForstoreRequest() *TmallInventoryQueryForstoreRequest{
    return &TmallInventoryQueryForstoreRequest{
        Params: model.NewParams(),
    }
}

func (r TmallInventoryQueryForstoreRequest) GetApiMethodName() string {
    return "tmall.inventory.query.forstore"
}

func (r TmallInventoryQueryForstoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallInventoryQueryForstoreRequest) SetParamList(paramList []InventoryQueryForStoreRequest) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r TmallInventoryQueryForstoreRequest) GetParamList() []InventoryQueryForStoreRequest {
    return r.paramList
}

