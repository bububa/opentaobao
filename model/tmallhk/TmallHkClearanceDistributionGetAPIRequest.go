package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallHkClearanceDistributionGetAPIRequest
分销供应商获取清关材料 API请求
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。 */
type TmallHkClearanceDistributionGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 是否需要身份证图片，不需要可以缩短接口响应时间
	_needImage bool
}

// New
