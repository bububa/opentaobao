package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机价格修改白名单 APIRequest
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

func NewAlibabaRetailVendingPriceWhitelistAddRequest() *AlibabaRetailVendingPriceWhitelistAddRequest{
    return &AlibabaRetailVendingPriceWhitelistAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.price.whitelist.add"
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidStarts(validStarts string) error {
    r.validStarts = validStarts
    r.Set("valid_starts", validStarts)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidStarts() string {
    return r.validStarts
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceCodeList(deviceCodeList []string) error {
    r.deviceCodeList = deviceCodeList
    r.Set("device_code_list", deviceCodeList)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceCodeList() []string {
    return r.deviceCodeList
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetDeviceUuidList(deviceUuidList []string) error {
    r.deviceUuidList = deviceUuidList
    r.Set("device_uuid_list", deviceUuidList)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetDeviceUuidList() []string {
    return r.deviceUuidList
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetValidEnds(validEnds string) error {
    r.validEnds = validEnds
    r.Set("valid_ends", validEnds)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetValidEnds() string {
    return r.validEnds
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetBarcode() string {
    return r.barcode
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetMinPrice(minPrice string) error {
    r.minPrice = minPrice
    r.Set("min_price", minPrice)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetMinPrice() string {
    return r.minPrice
}

func (r *AlibabaRetailVendingPriceWhitelistAddRequest) SetAllDevice(allDevice bool) error {
    r.allDevice = allDevice
    r.Set("all_device", allDevice)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistAddRequest) GetAllDevice() bool {
    return r.allDevice
}

