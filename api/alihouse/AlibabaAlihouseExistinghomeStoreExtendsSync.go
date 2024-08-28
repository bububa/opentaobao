package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreExtendsSync 门店扩展信息变更
// alibaba.alihouse.existinghome.store.extends.sync
//
// 门店扩展信息变更
func AlibabaAlihouseExistinghomeStoreExtendsSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
