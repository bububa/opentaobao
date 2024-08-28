package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalStoreAuditQuery ISV查询门店审核状态
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
func AlibabaAlihealthDentalStoreAuditQuery(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreAuditQueryAPIRequest, resp *alihealth2.AlibabaAlihealthDentalStoreAuditQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
