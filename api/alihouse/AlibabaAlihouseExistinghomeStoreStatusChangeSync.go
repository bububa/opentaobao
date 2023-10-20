package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreStatusChangeSync 门店状态变更
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
func AlibabaAlihouseExistinghomeStoreStatusChangeSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
