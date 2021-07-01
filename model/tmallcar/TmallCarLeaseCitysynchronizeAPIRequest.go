package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseCitysynchronizeAPIRequest
天猫开新车租后分期城市信息同步 API请求
tmall.car.lease.citysynchronize

天猫开新车租后分期城市信息同步 */
type TmallCarLeaseCitysynchronizeAPIRequest struct {
	model.Params
	// 地址信息
	_paramAreaDto *AreaDto
}

// New
