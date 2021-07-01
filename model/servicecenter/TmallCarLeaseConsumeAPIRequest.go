package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseConsumeAPIRequest
汽车租赁核销 API请求
tmall.car.lease.consume

租赁公司回传信息，核销 */
type TmallCarLeaseConsumeAPIRequest struct {
	model.Params
	// 核销请求
	_cosumeCodeReqDTO *CosumeCodeReqDto
}

// New
