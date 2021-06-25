package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库覆盖范围 APIRequest
taobao.region.warehouse.query

查询仓库覆盖范围
*/
type TaobaoRegionWarehouseQueryRequest struct {
    model.Params

    // 仓库编码
    storeCode   string 

}

func NewTaobaoRegionWarehouseQueryRequest() *TaobaoRegionWarehouseQueryRequest{
    return &TaobaoRegionWarehouseQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRegionWarehouseQueryRequest) GetApiMethodName() string {
    return "taobao.region.warehouse.query"
}

func (r TaobaoRegionWarehouseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRegionWarehouseQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoRegionWarehouseQueryRequest) GetStoreCode() string {
    return r.storeCode
}

