package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBankcardSync 结算账号同步
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
func AlibabaAlihouseExistinghomeBankcardSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
