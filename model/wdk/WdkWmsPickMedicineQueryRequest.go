package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询拣货单中的药品信息 APIRequest
wdk.wms.pick.medicine.query

联营商药机查询拣货单中的药品信息
*/
type WdkWmsPickMedicineQueryRequest struct {
    model.Params

    // shopId
    shopId   int64 

    // uuid
    uuid   string 

}

func NewWdkWmsPickMedicineQueryRequest() *WdkWmsPickMedicineQueryRequest{
    return &WdkWmsPickMedicineQueryRequest{
        Params: model.NewParams(),
    }
}

func (r WdkWmsPickMedicineQueryRequest) GetApiMethodName() string {
    return "wdk.wms.pick.medicine.query"
}

func (r WdkWmsPickMedicineQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkWmsPickMedicineQueryRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r WdkWmsPickMedicineQueryRequest) GetShopId() int64 {
    return r.shopId
}

func (r *WdkWmsPickMedicineQueryRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r WdkWmsPickMedicineQueryRequest) GetUuid() string {
    return r.uuid
}

