package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品编辑 API请求
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

// 初始化TmallItemStoreSchemaUpdateRequest对象
func NewTmallItemStoreSchemaUpdateRequest() *TmallItemStoreSchemaUpdateRequest{
    return &TmallItemStoreSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemStoreSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.store.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemStoreSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainItemId Setter
// 主商品ID
func (r *TmallItemStoreSchemaUpdateRequest) SetMainItemId(mainItemId int64) error {
    r.mainItemId = mainItemId
    r.Set("main_item_id", mainItemId)
    return nil
}

// MainItemId Getter
func (r TmallItemStoreSchemaUpdateRequest) GetMainItemId() int64 {
    return r.mainItemId
}
// StoreId Setter
// 门店ID
func (r *TmallItemStoreSchemaUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TmallItemStoreSchemaUpdateRequest) GetStoreId() int64 {
    return r.storeId
}
// Xml Setter
// 商品的schema xml
func (r *TmallItemStoreSchemaUpdateRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

// Xml Getter
func (r TmallItemStoreSchemaUpdateRequest) GetXml() string {
    return r.xml
}
