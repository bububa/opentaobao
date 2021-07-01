package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联营商药品柜核销 API请求
wdk.wms.pick.medicine.checksell

联营商药品柜核销
*/
type WdkWmsPickMedicineChecksellAPIRequest struct {
    model.Params
    // 从二维码扫描出的信息
    _uuid   string
    // shopId
    _shopId   int64
}

// 初始化WdkWmsPickMedicineChecksellAPIRequest对象
func NewWdkWmsPickMedicineChecksellRequest() *WdkWmsPickMedicineChecksellAPIRequest{
    return &WdkWmsPickMedicineChecksellAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.checksell"
}

// IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 从二维码扫描出的信息
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetUuid() string {
    return r._uuid
}
// ShopId Setter
// shopId
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetShopId(_shopId int64) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetShopId() int64 {
    return r._shopId
}
