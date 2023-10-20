package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomestorestatuschangesync 门店状态变更
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
func Alibabaalihouseexistinghomestorestatuschangesync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomestorestatuschangesyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomestorestatuschangesyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomestorestatuschangesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
