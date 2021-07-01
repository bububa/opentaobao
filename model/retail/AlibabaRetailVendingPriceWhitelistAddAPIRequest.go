package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailVendingPriceWhitelistAddAPIRequest
贩卖机价格修改白名单 API请求
alibaba.retail.vending.price.whitelist.add

贩卖机价格修改白名单 */
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

// New
