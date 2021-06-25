package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品修改规则获取 APIRequest
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
type TmallItemStoreUpdateSchemaGetRequest struct {
    model.Params

    // 主商品ID
    mainItemId   int64 

    // 门店ID
    storeId   int64 

}

func NewTmallItemStoreUpdateSchemaGetRequest() *TmallItemStoreUpdateSchemaGetRequest{
    return &TmallItemStoreUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemStoreUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.store.update.schema.get"
}

func (r TmallItemStoreUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemStoreUpdateSchemaGetRequest) SetMainItemId(mainItemId int64) error {
    r.mainItemId = mainItemId
    r.Set("main_item_id", mainItemId)
    return nil
}

func (r TmallItemStoreUpdateSchemaGetRequest) GetMainItemId() int64 {
    return r.mainItemId
}

func (r *TmallItemStoreUpdateSchemaGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TmallItemStoreUpdateSchemaGetRequest) GetStoreId() int64 {
    return r.storeId
}

