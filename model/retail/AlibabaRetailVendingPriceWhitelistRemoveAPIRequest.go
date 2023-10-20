package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailvendingpricewhitelistremoveAPIRequest 价格管控白名单去除 API请求
// alibaba.retail.vending.price.whitelist.remove
//
// 商家价格管控白名单去除
type AlibabaretailvendingpricewhitelistremoveAPIRequest struct {
	model.Params
	// 设备编码 device_code_list, device_uuid_list 二选一必填
	_deviceCodeList []string
	// 外部设备编码 device_code_list, device_uuid_list 二选一必填
	_deviceUuidList []string
	// 条码
	_barcode string
	// 淘宝用户ID
	_sellerId int64
	// 如果该参数传入，条码以商品条码为准
	_itemId int64
	// 是否生效到所有设备
	_allDevice bool
}

// NewAlibabaretailvendingpricewhitelistremoveRequest 初始化AlibabaretailvendingpricewhitelistremoveAPIRequest对象
func NewAlibabaretailvendingpricewhitelistremoveRequest() *AlibabaretailvendingpricewhitelistremoveAPIRequest {
	return &AlibabaretailvendingpricewhitelistremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.vending.price.whitelist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCodeList is DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetDeviceCodeList(_deviceCodeList []string) error {
	r._deviceCodeList = _deviceCodeList
	r.Set("device_code_list", _deviceCodeList)
	return nil
}

// GetDeviceCodeList DeviceCodeList Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetDeviceCodeList() []string {
	return r._deviceCodeList
}

// SetDeviceUuidList is DeviceUuidList Setter
// 外部设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetDeviceUuidList(_deviceUuidList []string) error {
	r._deviceUuidList = _deviceUuidList
	r.Set("device_uuid_list", _deviceUuidList)
	return nil
}

// GetDeviceUuidList DeviceUuidList Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetDeviceUuidList() []string {
	return r._deviceUuidList
}

// SetBarcode is Barcode Setter
// 条码
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSellerId is SellerId Setter
// 淘宝用户ID
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetItemId is ItemId Setter
// 如果该参数传入，条码以商品条码为准
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetAllDevice is AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaretailvendingpricewhitelistremoveAPIRequest) SetAllDevice(_allDevice bool) error {
	r._allDevice = _allDevice
	r.Set("all_device", _allDevice)
	return nil
}

// GetAllDevice AllDevice Getter
func (r AlibabaretailvendingpricewhitelistremoveAPIRequest) GetAllDevice() bool {
	return r._allDevice
}
