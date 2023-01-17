package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreBailSync 门店保证金余额同步
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
func AlibabaAlihouseExistinghomeStoreBailSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
