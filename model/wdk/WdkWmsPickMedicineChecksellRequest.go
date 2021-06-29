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
type WdkWmsPickMedicineChecksellRequest struct {
    model.Params
    // 从二维码扫描出的信息
    uuid   string
    // shopId
    shopId   int64
}

// 初始化WdkWmsPickMedicineChecksellRequest对象
func NewWdkWmsPickMedicineChecksellRequest() *WdkWmsPickMedicineChecksellRequest{
    return &WdkWmsPickMedicineChecksellRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineChecksellRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.checksell"
}

// IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineChecksellRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 从二维码扫描出的信息
func (r *WdkWmsPickMedicineChecksellRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r WdkWmsPickMedicineChecksellRequest) GetUuid() string {
    return r.uuid
}
// ShopId Setter
// shopId
func (r *WdkWmsPickMedicineChecksellRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r WdkWmsPickMedicineChecksellRequest) GetShopId() int64 {
    return r.shopId
}
