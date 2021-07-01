package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayAppinfoGetAPIRequest
tv支付获取应用信息 API请求
taobao.tvpay.appinfo.get

tv支付获取应用信息 */
type TaobaoTvpayAppinfoGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 客户端版本号
	_clientVersion string
}

// New
