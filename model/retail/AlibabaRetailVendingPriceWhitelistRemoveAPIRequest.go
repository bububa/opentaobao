package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailVendingPriceWhitelistRemoveAPIRequest
价格管控白名单去除 API请求
alibaba.retail.vending.price.whitelist.remove

商家价格管控白名单去除 */
type AlibabaRetailVendingPriceWhitelistRemoveAPIRequest struct {
	model.Params
	// 淘宝用户ID
	_sellerId int64
	// 设备编码 device_code_list, device_uuid_list 二选一必填
	_deviceCodeList []string
	// 外部设备编码 device_code_list, device_uuid_list 二选一必填
	_deviceUuidList []string
	// 条码
	_barcode string
	// 如果该参数传入，条码以商品条码为准
	_itemId int64
	// 是否生效到所有设备
	_allDevice bool
}

// NewAlibabaRetailVendingPriceWhitelistRemoveRequest 初始化AlibabaRetailVendingPriceWhitelistRemoveAPIRequest对象
func NewAlibabaRetailVendingPriceWhitelistRemoveRequest() *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest {
	return &AlibabaRetailVendingPriceWhitelistRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.vending.price.whitelist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// Set is DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetDeviceCodeList(_deviceCodeList []string) error {
	r._deviceCodeList = _deviceCodeList
	r.Set("device_code_list", _deviceCodeList)
	return nil
}

// Get DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetDeviceCodeList() []string {
	return r._deviceCodeList
}

// Set is DeviceUuidList Setter
// 外部设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetDeviceUuidList(_deviceUuidList []string) error {
	r._deviceUuidList = _deviceUuidList
	r.Set("device_uuid_list", _deviceUuidList)
	return nil
}

// Get DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetDeviceUuidList() []string {
	return r._deviceUuidList
}

// Set is Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// Get Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetBarcode() string {
	return r._barcode
}

// Set is ItemId Setter
// 如果该参数传入，条码以商品条码为准
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetAllDevice(_allDevice bool) error {
	r._allDevice = _allDevice
	r.Set("all_device", _allDevice)
	return nil
}

// Get AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetAllDevice() bool {
	return r._allDevice
}
