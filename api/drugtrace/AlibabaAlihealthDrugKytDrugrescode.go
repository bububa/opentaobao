package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.drugrescode
//
// 查询药品码段信息
func AlibabaAlihealthDrugKytDrugrescode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
