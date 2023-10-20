package healthnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/healthnr"
)

// Alibabahealthnrlogisticsquery 阿里健康新零售物流详情接口
// alibaba.health.nr.logistics.query
//
// 对阿里健康o2o对接的商户提供查询物流单详情的能力
func Alibabahealthnrlogisticsquery(clt *core.SDKClient, req *healthnr.AlibabahealthnrlogisticsqueryAPIRequest, session string) (*healthnr.AlibabahealthnrlogisticsqueryAPIResponse, error) {
	var resp healthnr.AlibabahealthnrlogisticsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
