package alidoc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthAlidocDrugStoreUpdate 更新药店
// alibaba.alihealth.alidoc.drug.store.update
//
// 药店信息更新接口
func AlibabaAlihealthAlidocDrugStoreUpdate(ctx context.Context, clt *core.SDKClient, req *alidoc.AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest, resp *alidoc.AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
