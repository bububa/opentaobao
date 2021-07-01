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
type AlibabaInteractRetailSaveshelflocationAPIRequest struct {
    model.Params
    // 门店code
    _storeCode   string
    // 货架编号
    _shelfNo   string
    // 经度
    _lng   string
    // 纬度
    _lat   string
    // POI
    _poiId   string
    // 地址
    _address   string
}

// 初始化AlibabaInteractRetailSaveshelflocationAPIRequest对象
func NewAlibabaInteractRetailSaveshelflocationRequest() *AlibabaInteractRetailSaveshelflocationAPIRequest{
    return &AlibabaInteractRetailSaveshelflocationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.saveshelflocation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 门店code
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetStoreCode() string {
    return r._storeCode
}
// ShelfNo Setter
// 货架编号
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetShelfNo(_shelfNo string) error {
    r._shelfNo = _shelfNo
    r.Set("shelf_no", _shelfNo)
    return nil
}

// ShelfNo Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetShelfNo() string {
    return r._shelfNo
}
// Lng Setter
// 经度
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetLng(_lng string) error {
    r._lng = _lng
    r.Set("lng", _lng)
    return nil
}

// Lng Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetLng() string {
    return r._lng
}
// Lat Setter
// 纬度
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetLat() string {
    return r._lat
}
// PoiId Setter
// POI
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetPoiId(_poiId string) error {
    r._poiId = _poiId
    r.Set("poi_id", _poiId)
    return nil
}

// PoiId Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetPoiId() string {
    return r._poiId
}
// Address Setter
// 地址
func (r *AlibabaInteractRetailSaveshelflocationAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaInteractRetailSaveshelflocationAPIRequest) GetAddress() string {
    return r._address
}
