package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugdetail 查询药品详细信息
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
func AlibabaAlihealthDrugKytDrugdetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugdetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
