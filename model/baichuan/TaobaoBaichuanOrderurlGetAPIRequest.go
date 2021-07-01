package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOrderurlGetAPIRequest
百川订单详情 API请求
taobao.baichuan.orderurl.get

百川订单详情 */
type TaobaoBaichuanOrderurlGetAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
