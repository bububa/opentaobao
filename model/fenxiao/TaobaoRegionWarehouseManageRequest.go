package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑仓库覆盖范围 API请求
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
type TaobaoRegionWarehouseManageRequest struct {
    model.Params
    // 仓库编码
    _storeCode   string
    // 可映射三级地址,例: 广东省
    _regions   []string
}

// 初始化TaobaoRegionWarehouseManageRequest对象
func NewTaobaoRegionWarehouseManageRequest() *TaobaoRegionWarehouseManageRequest{
    return &TaobaoRegionWarehouseManageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionWarehouseManageRequest) GetApiMethodName() string {
    return "taobao.region.warehouse.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionWarehouseManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoRegionWarehouseManageRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoRegionWarehouseManageRequest) GetStoreCode() string {
    return r._storeCode
}
// Regions Setter
// 可映射三级地址,例: 广东省
func (r *TaobaoRegionWarehouseManageRequest) SetRegions(_regions []string) error {
    r._regions = _regions
    r.Set("regions", _regions)
    return nil
}

// Regions Getter
func (r TaobaoRegionWarehouseManageRequest) GetRegions() []string {
    return r._regions
}
