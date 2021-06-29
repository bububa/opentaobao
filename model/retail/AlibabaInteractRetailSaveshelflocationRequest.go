package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存地理位置和货架关系 APIRequest
alibaba.interact.retail.saveshelflocation

保存地理位置和货架关系
*/
type AlibabaInteractRetailSaveshelflocationRequest struct {
    model.Params

    // 门店code
    storeCode   string 

    // 货架编号
    shelfNo   string 

    // 经度
    lng   string 

    // 纬度
    lat   string 

    // POI
    poiId   string 

    // 地址
    address   string 

}

func NewAlibabaInteractRetailSaveshelflocationRequest() *AlibabaInteractRetailSaveshelflocationRequest{
    return &AlibabaInteractRetailSaveshelflocationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.saveshelflocation"
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractRetailSaveshelflocationRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *AlibabaInteractRetailSaveshelflocationRequest) SetShelfNo(shelfNo string) error {
    r.shelfNo = shelfNo
    r.Set("shelf_no", shelfNo)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetShelfNo() string {
    return r.shelfNo
}

func (r *AlibabaInteractRetailSaveshelflocationRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetLng() string {
    return r.lng
}

func (r *AlibabaInteractRetailSaveshelflocationRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetLat() string {
    return r.lat
}

func (r *AlibabaInteractRetailSaveshelflocationRequest) SetPoiId(poiId string) error {
    r.poiId = poiId
    r.Set("poi_id", poiId)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetPoiId() string {
    return r.poiId
}

func (r *AlibabaInteractRetailSaveshelflocationRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaInteractRetailSaveshelflocationRequest) GetAddress() string {
    return r.address
}

