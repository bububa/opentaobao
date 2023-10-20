package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreSync 二手房标准门店数据同步
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
func AlibabaAlihouseExistinghomeStoreSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeStoreSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeStoreSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
