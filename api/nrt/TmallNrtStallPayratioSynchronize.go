package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStallPayratioSynchronize 同步摊位收银比例
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
func TmallNrtStallPayratioSynchronize(clt *core.SDKClient, req *nrt.TmallNrtStallPayratioSynchronizeAPIRequest, resp *nrt.TmallNrtStallPayratioSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
