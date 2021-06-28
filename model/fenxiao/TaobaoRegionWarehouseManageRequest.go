package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑仓库覆盖范围 APIRequest
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
type TaobaoRegionWarehouseManageRequest struct {
    model.Params

    // 仓库编码
    storeCode   string 

    // 可映射三级地址,例: 广东省
    regions   []string 

}

func NewTaobaoRegionWarehouseManageRequest() *TaobaoRegionWarehouseManageRequest{
    return &TaobaoRegionWarehouseManageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRegionWarehouseManageRequest) GetApiMethodName() string {
    return "taobao.region.warehouse.manage"
}

func (r TaobaoRegionWarehouseManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRegionWarehouseManageRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoRegionWarehouseManageRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoRegionWarehouseManageRequest) SetRegions(regions []string) error {
    r.regions = regions
    r.Set("regions", regions)
    return nil
}

func (r TaobaoRegionWarehouseManageRequest) GetRegions() []string {
    return r.regions
}

