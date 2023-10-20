package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStallSynchronize 摊位信息同步
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
func TmallNrtStallSynchronize(clt *core.SDKClient, req *nrt.TmallNrtStallSynchronizeAPIRequest, resp *nrt.TmallNrtStallSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
