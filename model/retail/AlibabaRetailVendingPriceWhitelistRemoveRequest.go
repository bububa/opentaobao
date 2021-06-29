package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格管控白名单去除 APIRequest
alibaba.retail.vending.price.whitelist.remove

商家价格管控白名单去除
*/
type AlibabaRetailVendingPriceWhitelistRemoveRequest struct {
    model.Params

    // 淘宝用户ID
    sellerId   int64 

    // 设备编码 device_code_list, device_uuid_list 二选一必填
    deviceCodeList   []string 

    // 外部设备编码 device_code_list, device_uuid_list 二选一必填
    deviceUuidList   []string 

    // 条码
    barcode   string 

    // 如果该参数传入，条码以商品条码为准
    itemId   int64 

    // 是否生效到所有设备
    allDevice   bool 

}

func NewAlibabaRetailVendingPriceWhitelistRemoveRequest() *AlibabaRetailVendingPriceWhitelistRemoveRequest{
    return &AlibabaRetailVendingPriceWhitelistRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetApiMethodName() string {
    return "alibaba.retail.vending.price.whitelist.remove"
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetDeviceCodeList(deviceCodeList []string) error {
    r.deviceCodeList = deviceCodeList
    r.Set("device_code_list", deviceCodeList)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetDeviceCodeList() []string {
    return r.deviceCodeList
}

func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetDeviceUuidList(deviceUuidList []string) error {
    r.deviceUuidList = deviceUuidList
    r.Set("device_uuid_list", deviceUuidList)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetDeviceUuidList() []string {
    return r.deviceUuidList
}

func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetBarcode() string {
    return r.barcode
}

func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaRetailVendingPriceWhitelistRemoveRequest) SetAllDevice(allDevice bool) error {
    r.allDevice = allDevice
    r.Set("all_device", allDevice)
    return nil
}

func (r AlibabaRetailVendingPriceWhitelistRemoveRequest) GetAllDevice() bool {
    return r.allDevice
}

