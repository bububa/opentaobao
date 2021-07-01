package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiDriverBlacklistAddAPIRequest
添加司机黑名单 API请求
alibaba.happytrip.taxi.driver.blacklist.add

实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行 */
type AlibabaHappytripTaxiDriverBlacklistAddAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// New
