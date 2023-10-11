package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreExtendsSync 门店扩展信息变更
// alibaba.alihouse.existinghome.store.extends.sync
//
// 门店扩展信息变更
func AlibabaAlihouseExistinghomeStoreExtendsSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
