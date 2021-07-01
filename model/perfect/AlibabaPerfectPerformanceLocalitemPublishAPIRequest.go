package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPerfectPerformanceLocalitemPublishAPIRequest
同城购定制化发品 API请求
alibaba.perfect.performance.localitem.publish

同城购业务定制化发品接口，同城购业务线专用 */
type AlibabaPerfectPerformanceLocalitemPublishAPIRequest struct {
	model.Params
	// 请求参数
	_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq
}

// New
