package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

/* AlibabaPerfectPerformanceLocalitemPublish
同城购定制化发品
alibaba.perfect.performance.localitem.publish

同城购业务定制化发品接口，同城购业务线专用 */
func AlibabaPerfectPerformanceLocalitemPublish(clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceLocalitemPublishAPIRequest, session string) (*perfect.AlibabaPerfectPerformanceLocalitemPublishAPIResponse, error) {
	var resp perfect.AlibabaPerfectPerformanceLocalitemPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
