package retail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailVendingPriceWhitelistRemoveAPIRequest 价格管控白名单去除 API请求
// alibaba.retail.vending.price.whitelist.remove
//
// 商家价格管控白名单去除
type AlibabaRetailVendingPriceWhitelistRemoveAPIRequest struct {
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

// NewAlibabaRetailVendingPriceWhitelistRemoveRequest 初始化AlibabaRetailVendingPriceWhitelistRemoveAPIRequest对象
func NewAlibabaRetailVendingPriceWhitelistRemoveRequest() *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest {
	return &AlibabaRetailVendingPriceWhitelistRemoveAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) Reset() {
	r._deviceCodeList = r._deviceCodeList[:0]
	r._deviceUuidList = r._deviceUuidList[:0]
	r._barcode = ""
	r._sellerId = 0
	r._itemId = 0
	r._allDevice = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.vending.price.whitelist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCodeList is DeviceCodeList Setter
// 设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetDeviceCodeList(_deviceCodeList []string) error {
	r._deviceCodeList = _deviceCodeList
	r.Set("device_code_list", _deviceCodeList)
	return nil
}

// GetDeviceCodeList DeviceCodeList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetDeviceCodeList() []string {
	return r._deviceCodeList
}

// SetDeviceUuidList is DeviceUuidList Setter
// 外部设备编码 device_code_list, device_uuid_list 二选一必填
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetDeviceUuidList(_deviceUuidList []string) error {
	r._deviceUuidList = _deviceUuidList
	r.Set("device_uuid_list", _deviceUuidList)
	return nil
}

// GetDeviceUuidList DeviceUuidList Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetDeviceUuidList() []string {
	return r._deviceUuidList
}

// SetBarcode is Barcode Setter
// 条码
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSellerId is SellerId Setter
// 淘宝用户ID
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetItemId is ItemId Setter
// 如果该参数传入，条码以商品条码为准
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetAllDevice is AllDevice Setter
// 是否生效到所有设备
func (r *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) SetAllDevice(_allDevice bool) error {
	r._allDevice = _allDevice
	r.Set("all_device", _allDevice)
	return nil
}

// GetAllDevice AllDevice Getter
func (r AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) GetAllDevice() bool {
	return r._allDevice
}

var poolAlibabaRetailVendingPriceWhitelistRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailVendingPriceWhitelistRemoveRequest()
	},
}

// GetAlibabaRetailVendingPriceWhitelistRemoveRequest 从 sync.Pool 获取 AlibabaRetailVendingPriceWhitelistRemoveAPIRequest
func GetAlibabaRetailVendingPriceWhitelistRemoveAPIRequest() *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest {
	return poolAlibabaRetailVendingPriceWhitelistRemoveAPIRequest.Get().(*AlibabaRetailVendingPriceWhitelistRemoveAPIRequest)
}

// ReleaseAlibabaRetailVendingPriceWhitelistRemoveAPIRequest 将 AlibabaRetailVendingPriceWhitelistRemoveAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailVendingPriceWhitelistRemoveAPIRequest(v *AlibabaRetailVendingPriceWhitelistRemoveAPIRequest) {
	v.Reset()
	poolAlibabaRetailVendingPriceWhitelistRemoveAPIRequest.Put(v)
}
