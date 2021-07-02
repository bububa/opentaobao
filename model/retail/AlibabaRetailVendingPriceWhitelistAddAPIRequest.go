package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailVendingPriceWhitelistAddAPIRequest 贩卖机价格修改白名单 API请求
// alibaba.retail.vending.price.whitelist.add
//
// 贩卖机价格修改白名单
type AlibabaRetailVendingPriceWhitelistAddAPIRequest struct {
	model.Params
	// 生效时间
	_validStarts string
	// 淘宝用户ID
	_sellerId int64
	// 设备编码 device_code_list, device_uuid_list 二选一必填
	_deviceCodeList []string
	// 外部设备编码   device_code_list, device_uuid_list 二选一必填
	_deviceUuidList []string
	// 生效结束时间
	_validEnds string
	// 条码
	_barcode string
	// 商品ID
	_itemId int64
	// 允许修改的最低价
	_minPrice string
	// 是否生效到所有设备
	_allDevice bool
}

// NewAlibabaRetailVendingPriceWhitelistAddRequest 初始化AlibabaRetailVendingPriceWhitelistAddAPIRequest对象
func NewAlibabaRetailVendingPriceWhitelistAddRequest() *AlibabaRetailVendingPriceWhitelistAddAPIRequest {
	return &AlibabaRetailVendingPriceWhitelistAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.vending.price.whitelist.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ValidStarts Setter
// 生效时间
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetValidStarts(_validStarts string) error {
	r._validStarts = _validStarts
	r.Set("valid_starts", _validStarts)
	return nil
}

// Get ValidStarts Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetValidStarts() string {
	return r._validStarts
}

// Set is SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// Set is DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetDeviceCodeList(_deviceCodeList []string) error {
	r._deviceCodeList = _deviceCodeList
	r.Set("device_code_list", _deviceCodeList)
	return nil
}

// Get DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetDeviceCodeList() []string {
	return r._deviceCodeList
}

// Set is DeviceUuidList Setter
// 外部设备编码   device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetDeviceUuidList(_deviceUuidList []string) error {
	r._deviceUuidList = _deviceUuidList
	r.Set("device_uuid_list", _deviceUuidList)
	return nil
}

// Get DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetDeviceUuidList() []string {
	return r._deviceUuidList
}

// Set is ValidEnds Setter
// 生效结束时间
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetValidEnds(_validEnds string) error {
	r._validEnds = _validEnds
	r.Set("valid_ends", _validEnds)
	return nil
}

// Get ValidEnds Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetValidEnds() string {
	return r._validEnds
}

// Set is Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// Get Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetBarcode() string {
	return r._barcode
}

// Set is ItemId Setter
// 商品ID
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is MinPrice Setter
// 允许修改的最低价
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetMinPrice(_minPrice string) error {
	r._minPrice = _minPrice
	r.Set("min_price", _minPrice)
	return nil
}

// Get MinPrice Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetMinPrice() string {
	return r._minPrice
}

// Set is AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistAddAPIRequest) SetAllDevice(_allDevice bool) error {
	r._allDevice = _allDevice
	r.Set("all_device", _allDevice)
	return nil
}

// Get AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistAddAPIRequest) GetAllDevice() bool {
	return r._allDevice
}
