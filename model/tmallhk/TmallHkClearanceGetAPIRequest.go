package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallHkClearanceGetAPIRequest
天猫国际-清关材料查询 API请求
tmall.hk.clearance.get

提供订单收货人身份信息查询功能。 */
type TmallHkClearanceGetAPIRequest struct {
	model.Params
	// 天猫国际订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// New
