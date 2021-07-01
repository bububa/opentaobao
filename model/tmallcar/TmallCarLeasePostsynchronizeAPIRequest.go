package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeasePostsynchronizeAPIRequest
天猫开新车租后信息同步 API请求
tmall.car.lease.postsynchronize

商家同步天猫开新车租后方案 */
type TmallCarLeasePostsynchronizeAPIRequest struct {
	model.Params
	// 租后方案信息
	_schemeDto *CarLeasePostSchemeSynchronizeDto
}

// New
