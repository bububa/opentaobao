package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库覆盖范围 API请求
taobao.region.warehouse.query

查询仓库覆盖范围
*/
type TaobaoRegionWarehouseQueryRequest struct {
    model.Params
    // 仓库编码
    _storeCode   string
}

// 初始化TaobaoRegionWarehouseQueryRequest对象
func NewTaobaoRegionWarehouseQueryRequest() *TaobaoRegionWarehouseQueryRequest{
    return &TaobaoRegionWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionWarehouseQueryRequest) GetApiMethodName() string {
    return "taobao.region.warehouse.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoRegionWarehouseQueryRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoRegionWarehouseQueryRequest) GetStoreCode() string {
    return r._storeCode
}
