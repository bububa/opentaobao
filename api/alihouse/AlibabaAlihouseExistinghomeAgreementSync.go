package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeagreementsync 二手房电子协议数据同步
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
func Alibabaalihouseexistinghomeagreementsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeagreementsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeagreementsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeagreementsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
