package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalBindAuditQuery ISV查询绑定审核状态
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
func AlibabaAlihealthDentalBindAuditQuery(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalBindAuditQueryAPIRequest, resp *alihealth2.AlibabaAlihealthDentalBindAuditQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
