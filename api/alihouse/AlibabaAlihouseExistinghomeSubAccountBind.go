package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomesubaccountbind 子账号入驻
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
func Alibabaalihouseexistinghomesubaccountbind(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomesubaccountbindAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomesubaccountbindAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomesubaccountbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
