package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebankcardsync 结算账号同步
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
func Alibabaalihouseexistinghomebankcardsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebankcardsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebankcardsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebankcardsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
