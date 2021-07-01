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

// New
