package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderCancel 大麦-库存释放
// alibaba.damai.maitix.order.cancel
//
// 库存释放
func AlibabaDamaiMaitixOrderCancel(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderCancelAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
