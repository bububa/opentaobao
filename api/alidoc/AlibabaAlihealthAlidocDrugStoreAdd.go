package alidoc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthAlidocDrugStoreAdd gsk新增药店
// alibaba.alihealth.alidoc.drug.store.add
//
// GSK上传药店信息
func AlibabaAlihealthAlidocDrugStoreAdd(ctx context.Context, clt *core.SDKClient, req *alidoc.AlibabaAlihealthAlidocDrugStoreAddAPIRequest, resp *alidoc.AlibabaAlihealthAlidocDrugStoreAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
