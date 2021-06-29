package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库信息 API请求
taobao.inventory.store.query

查询商家仓信息
*/
type TaobaoInventoryStoreQueryRequest struct {
    model.Params
    // 商家的仓库编码
    _storeCode   string
}

// 初始化TaobaoInventoryStoreQueryRequest对象
func NewTaobaoInventoryStoreQueryRequest() *TaobaoInventoryStoreQueryRequest{
    return &TaobaoInventoryStoreQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryStoreQueryRequest) GetApiMethodName() string {
    return "taobao.inventory.store.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 商家的仓库编码
func (r *TaobaoInventoryStoreQueryRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoInventoryStoreQueryRequest) GetStoreCode() string {
    return r._storeCode
}
