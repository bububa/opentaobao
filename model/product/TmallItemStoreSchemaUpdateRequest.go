package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品编辑 APIRequest
tmall.item.store.schema.update

天猫门店商品编辑
*/
type TmallItemStoreSchemaUpdateRequest struct {
    model.Params

    // 主商品ID
    mainItemId   int64 

    // 门店ID
    storeId   int64 

    // 商品的schema xml
    xml   string 

}

func NewTmallItemStoreSchemaUpdateRequest() *TmallItemStoreSchemaUpdateRequest{
    return &TmallItemStoreSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemStoreSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.store.schema.update"
}

func (r TmallItemStoreSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemStoreSchemaUpdateRequest) SetMainItemId(mainItemId int64) error {
    r.mainItemId = mainItemId
    r.Set("main_item_id", mainItemId)
    return nil
}

func (r TmallItemStoreSchemaUpdateRequest) GetMainItemId() int64 {
    return r.mainItemId
}

func (r *TmallItemStoreSchemaUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TmallItemStoreSchemaUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TmallItemStoreSchemaUpdateRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

func (r TmallItemStoreSchemaUpdateRequest) GetXml() string {
    return r.xml
}

