package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest
移除司机黑名单 API请求
alibaba.happytrip.taxi.driver.blacklist.remove

移除司机黑名单 */
type AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// New
