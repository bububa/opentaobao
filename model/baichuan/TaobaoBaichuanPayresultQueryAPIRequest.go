package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanPayresultQueryAPIRequest
百川支付完成回调 API请求
taobao.baichuan.payresult.query

百川支付完成回调 */
type TaobaoBaichuanPayresultQueryAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
