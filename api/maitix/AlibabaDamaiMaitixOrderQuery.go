package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixorderquery 大麦-查询分销单
// alibaba.damai.maitix.order.query
//
// 查询分销单
func Alibabadamaimaitixorderquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixorderqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixorderqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
