package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseRiskcallbackAPIRequest
整车租赁风控模型回调 API请求
tmall.car.lease.riskcallback

租赁公司回调风控结果 */
type TmallCarLeaseRiskcallbackAPIRequest struct {
	model.Params
	// 授信结果
	_creditInfo *CreditInfoTopDto
}

// New
