package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPerfectPerformanceLocalitemEditAPIRequest
同城购定制发品编辑 API请求
alibaba.perfect.performance.localitem.edit

同城购业务定制化发品接口，同城购业务线专用 */
type AlibabaPerfectPerformanceLocalitemEditAPIRequest struct {
	model.Params
	// 请求参数
	_paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq
}

// New
