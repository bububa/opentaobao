package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreBailSync 门店保证金余额同步
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
func AlibabaAlihouseExistinghomeStoreBailSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
