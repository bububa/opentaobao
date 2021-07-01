package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫门店商品修改规则获取 API请求
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
type TmallItemStoreUpdateSchemaGetAPIRequest struct {
    model.Params
    // 主商品ID
    _mainItemId   int64
    // 门店ID
    _storeId   int64
}

// 初始化TmallItemStoreUpdateSchemaGetAPIRequest对象
func NewTmallItemStoreUpdateSchemaGetRequest() *TmallItemStoreUpdateSchemaGetAPIRequest{
    return &TmallItemStoreUpdateSchemaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetApiMethodName() string {
    return "tmall.item.store.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainItemId Setter
// 主商品ID
func (r *TmallItemStoreUpdateSchemaGetAPIRequest) SetMainItemId(_mainItemId int64) error {
    r._mainItemId = _mainItemId
    r.Set("main_item_id", _mainItemId)
    return nil
}

// MainItemId Getter
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetMainItemId() int64 {
    return r._mainItemId
}
// StoreId Setter
// 门店ID
func (r *TmallItemStoreUpdateSchemaGetAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TmallItemStoreUpdateSchemaGetAPIRequest) GetStoreId() int64 {
    return r._storeId
}
