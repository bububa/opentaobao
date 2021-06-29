package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询拣货单中的药品信息 API请求
wdk.wms.pick.medicine.query

联营商药机查询拣货单中的药品信息
*/
type WdkWmsPickMedicineQueryRequest struct {
    model.Params
    // shopId
    _shopId   int64
    // uuid
    _uuid   string
}

// 初始化WdkWmsPickMedicineQueryRequest对象
func NewWdkWmsPickMedicineQueryRequest() *WdkWmsPickMedicineQueryRequest{
    return &WdkWmsPickMedicineQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineQueryRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.query"
}

// IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// shopId
func (r *WdkWmsPickMedicineQueryRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r WdkWmsPickMedicineQueryRequest) GetShopId() int64 {
    return r._shopId
}
// Uuid Setter
// uuid
func (r *WdkWmsPickMedicineQueryRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r WdkWmsPickMedicineQueryRequest) GetUuid() string {
    return r._uuid
}
