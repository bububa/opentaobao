package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// AlibabaPerfectPerformanceLocalitemEdit 同城购定制发品编辑
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
func AlibabaPerfectPerformanceLocalitemEdit(clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceLocalitemEditAPIRequest, resp *perfect.AlibabaPerfectPerformanceLocalitemEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
