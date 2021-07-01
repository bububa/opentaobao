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
type WdkWmsPickMedicineQueryAPIRequest struct {
    model.Params
    // shopId
    _shopId   int64
    // uuid
    _uuid   string
}

// 初始化WdkWmsPickMedicineQueryAPIRequest对象
func NewWdkWmsPickMedicineQueryRequest() *WdkWmsPickMedicineQueryAPIRequest{
    return &WdkWmsPickMedicineQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineQueryAPIRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.query"
}

// IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// shopId
func (r *WdkWmsPickMedicineQueryAPIRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r WdkWmsPickMedicineQueryAPIRequest) GetShopId() int64 {
    return r._shopId
}
// Uuid Setter
// uuid
func (r *WdkWmsPickMedicineQueryAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r WdkWmsPickMedicineQueryAPIRequest) GetUuid() string {
    return r._uuid
}
