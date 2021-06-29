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
    validStarts   string
    // 淘宝用户ID
    sellerId   int64
    // 设备编码 device_code_list, device_uuid_list 二选一必填
    deviceCodeList   []string
    // 外部设备编码   device_code_list, device_uuid_list 二选一必填
    deviceUuidList   []string
    // 生效结束时间
    validEnds   string
    // 条码
    barcode   string
    // 商品ID
    itemId   int64
    // 允许修改的最低价
    minPrice   string
    // 是否生效到所有设备
    allDevice   bool
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
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidStarts(validStarts string) error {
    r.validStarts = validStarts
    r.Set("valid_starts", validStarts)
    return nil
}

// ValidStarts Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidStarts() string {
    return r.validStarts
}
// SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetSellerId() int64 {
    return r.sellerId
}
// DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceCodeList(deviceCodeList []string) error {
    r.deviceCodeList = deviceCodeList
    r.Set("device_code_list", deviceCodeList)
    return nil
}

// DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceCodeList() []string {
    return r.deviceCodeList
}
// DeviceUuidList Setter
// 外部设备编码   device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceUuidList(deviceUuidList []string) error {
    r.deviceUuidList = deviceUuidList
    r.Set("device_uuid_list", deviceUuidList)
    return nil
}

// DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceUuidList() []string {
    return r.deviceUuidList
}
// ValidEnds Setter
// 生效结束时间
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidEnds(validEnds string) error {
    r.validEnds = validEnds
    r.Set("valid_ends", validEnds)
    return nil
}

// ValidEnds Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidEnds() string {
    return r.validEnds
}
// Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetBarcode() string {
    return r.barcode
}
// ItemId Setter
// 商品ID
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetItemId() int64 {
    return r.itemId
}
// MinPrice Setter
// 允许修改的最低价
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetMinPrice(minPrice string) error {
    r.minPrice = minPrice
    r.Set("min_price", minPrice)
    return nil
}

// MinPrice Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetMinPrice() string {
    return r.minPrice
}
// AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetAllDevice(allDevice bool) error {
    r.allDevice = allDevice
    r.Set("all_device", allDevice)
    return nil
}

// AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetAllDevice() bool {
    return r.allDevice
}
