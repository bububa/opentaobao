package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreSync 二手房标准门店数据同步
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
func AlibabaAlihouseExistinghomeStoreSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
