package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreStatusChangeSync 门店状态变更
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
func AlibabaAlihouseExistinghomeStoreStatusChangeSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
