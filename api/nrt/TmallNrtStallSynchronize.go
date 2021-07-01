package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

/* TmallNrtStallSynchronize
摊位信息同步
tmall.nrt.stall.synchronize

摊位信息同步 */
func TmallNrtStallSynchronize(clt *core.SDKClient, req *nrt.TmallNrtStallSynchronizeAPIRequest, session string) (*nrt.TmallNrtStallSynchronizeAPIResponse, error) {
	var resp nrt.TmallNrtStallSynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
