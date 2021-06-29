package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机价格修改白名单 API请求
alibaba.retail.vending.price.whitelist.add

贩卖机价格修改白名单
*/
type AlibabaRetailVendingPriceWhitelistAddRequest struct {
    model.Params
    // 生效时间
    _validStarts   string
    // 淘宝用户ID
    _sellerId   int64
    // 设备编码 device_code_list, device_uuid_list 二选一必填
    _deviceCodeList   []string
    // 外部设备编码   device_code_list, device_uuid_list 二选一必填
    _deviceUuidList   []string
    // 生效结束时间
    _validEnds   string
    // 条码
    _barcode   string
    // 商品ID
    _itemId   int64
    // 允许修改的最低价
    _minPrice   string
    // 是否生效到所有设备
    _allDevice   bool
}

// 初始化AlibabaRetailVendingPriceWhitelistAddRequest对象
func NewAlibabaRetailVendingPriceWhitelistAddRequest() *AlibabaRetailVendingPriceWhitelistAddRequest{
    return &AlibabaRetailVendingPriceWhitelistAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.price.whitelist.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ValidStarts Setter
// 生效时间
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidStarts(_validStarts string) error {
    r._validStarts = _validStarts
    r.Set("valid_starts", _validStarts)
    return nil
}

// ValidStarts Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidStarts() string {
    return r._validStarts
}
// SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetSellerId() int64 {
    return r._sellerId
}
// DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceCodeList(_deviceCodeList []string) error {
    r._deviceCodeList = _deviceCodeList
    r.Set("device_code_list", _deviceCodeList)
    return nil
}

// DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceCodeList() []string {
    return r._deviceCodeList
}
// DeviceUuidList Setter
// 外部设备编码   device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceUuidList(_deviceUuidList []string) error {
    r._deviceUuidList = _deviceUuidList
    r.Set("device_uuid_list", _deviceUuidList)
    return nil
}

// DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceUuidList() []string {
    return r._deviceUuidList
}
// ValidEnds Setter
// 生效结束时间
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidEnds(_validEnds string) error {
    r._validEnds = _validEnds
    r.Set("valid_ends", _validEnds)
    return nil
}

// ValidEnds Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidEnds() string {
    return r._validEnds
}
// Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetBarcode() string {
    return r._barcode
}
// ItemId Setter
// 商品ID
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetItemId() int64 {
    return r._itemId
}
// MinPrice Setter
// 允许修改的最低价
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetMinPrice(_minPrice string) error {
    r._minPrice = _minPrice
    r.Set("min_price", _minPrice)
    return nil
}

// MinPrice Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetMinPrice() string {
    return r._minPrice
}
// AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetAllDevice(_allDevice bool) error {
    r._allDevice = _allDevice
    r.Set("all_device", _allDevice)
    return nil
}

// AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetAllDevice() bool {
    return r._allDevice
}
