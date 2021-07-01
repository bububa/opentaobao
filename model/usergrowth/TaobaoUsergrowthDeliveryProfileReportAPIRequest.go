package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDeliveryProfileReportAPIRequest
标签上报 API请求
taobao.usergrowth.delivery.profile.report

渠道上报标签信息 */
type TaobaoUsergrowthDeliveryProfileReportAPIRequest struct {
	model.Params
	// 标签参数, 支持一次传多个， 一次最多传20个
	_data string
	// 渠道标识，找淘宝运营申请
	_channel string
}

// New
