package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanTaokeTraceAPIRequest
百川淘客打点 API请求
taobao.baichuan.taoke.trace

百川淘客打点 */
type TaobaoBaichuanTaokeTraceAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
