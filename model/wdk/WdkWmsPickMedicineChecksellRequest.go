package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联营商药品柜核销 APIRequest
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

func NewWdkWmsPickMedicineChecksellRequest() *WdkWmsPickMedicineChecksellRequest{
    return &WdkWmsPickMedicineChecksellRequest{
        Params: model.NewParams(),
    }
}

func (r WdkWmsPickMedicineChecksellRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.checksell"
}

func (r WdkWmsPickMedicineChecksellRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkWmsPickMedicineChecksellRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r WdkWmsPickMedicineChecksellRequest) GetUuid() string {
    return r.uuid
}

func (r *WdkWmsPickMedicineChecksellRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r WdkWmsPickMedicineChecksellRequest) GetShopId() int64 {
    return r.shopId
}

