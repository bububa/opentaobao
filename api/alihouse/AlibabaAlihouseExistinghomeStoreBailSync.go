package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomestorebailsync 门店保证金余额同步
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
func Alibabaalihouseexistinghomestorebailsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomestorebailsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomestorebailsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomestorebailsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
