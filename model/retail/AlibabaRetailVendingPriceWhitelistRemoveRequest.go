package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格管控白名单去除 API请求
alibaba.retail.vending.price.whitelist.remove

商家价格管控白名单去除
*/
type AlibabaRetailVendingPriceWhitelistRemoveRequest struct {
    model.Params
    // 淘宝用户ID
    _sellerId   int64
    // 设备编码 device_code_list, device_uuid_list 二选一必填
    _deviceCodeList   []string
    // 外部设备编码 device_code_list, device_uuid_list 二选一必填
    _deviceUuidList   []string
    // 条码
    _barcode   string
    // 如果该参数传入，条码以商品条码为准
    _itemId   int64
    // 是否生效到所有设备
    _allDevice   bool
}

// 初始化AlibabaRetailVendingPriceWhitelistRemoveRequest对象
func NewAlibabaRetailVendingPriceWhitelistRemoveRequest() *AlibabaRetailVendingPriceWhitelistRemoveRequest{
    return &AlibabaRetailVendingPriceWhitelistRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.price.whitelist.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetSellerId() int64 {
    return r._sellerId
}
// DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetDeviceCodeList(_deviceCodeList []string) error {
    r._deviceCodeList = _deviceCodeList
    r.Set("device_code_list", _deviceCodeList)
    return nil
}

// DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetDeviceCodeList() []string {
    return r._deviceCodeList
}
// DeviceUuidList Setter
// 外部设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetDeviceUuidList(_deviceUuidList []string) error {
    r._deviceUuidList = _deviceUuidList
    r.Set("device_uuid_list", _deviceUuidList)
    return nil
}

// DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetDeviceUuidList() []string {
    return r._deviceUuidList
}
// Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetBarcode() string {
    return r._barcode
}
// ItemId Setter
// 如果该参数传入，条码以商品条码为准
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetItemId() int64 {
    return r._itemId
}
// AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetAllDevice(_allDevice bool) error {
    r._allDevice = _allDevice
    r.Set("all_device", _allDevice)
    return nil
}

// AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetAllDevice() bool {
    return r._allDevice
}