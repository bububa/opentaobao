package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存地理位置和货架关系 API请求
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

// 初始化AlibabaInteractRetailSaveshelflocationRequest对象
func NewAlibabaInteractRetailSaveshelflocationRequest() *AlibabaInteractRetailSaveshelflocationRequest{
    return &AlibabaInteractRetailSaveshelflocationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractRetailSaveshelflocationRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.saveshelflocation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractRetailSaveshelflocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 门店code
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetStoreCode() string {
    return r.storeCode
}
// ShelfNo Setter
// 货架编号
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetShelfNo(shelfNo string) error {
    r.shelfNo = shelfNo
    r.Set("shelf_no", shelfNo)
    return nil
}

// ShelfNo Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetShelfNo() string {
    return r.shelfNo
}
// Lng Setter
// 经度
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

// Lng Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetLng() string {
    return r.lng
}
// Lat Setter
// 纬度
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetLat() string {
    return r.lat
}
// PoiId Setter
// POI
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetPoiId(poiId string) error {
    r.poiId = poiId
    r.Set("poi_id", poiId)
    return nil
}

// PoiId Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetPoiId() string {
    return r.poiId
}
// Address Setter
// 地址
func (r *AlibabaInteractRetailSaveshelflocationRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaInteractRetailSaveshelflocationRequest) GetAddress() string {
    return r.address
}
