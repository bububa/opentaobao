package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterAnomalyrecourseSearchAPIRequest
天猫服务平台服务商一键求助单查询 API请求
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询 */
type TmallServicecenterAnomalyrecourseSearchAPIRequest struct {
	model.Params
	// 开始时间
	_start int64
	// 结束时间
	_end int64
}

// New
