package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSaveent 新增往来单位企业
// alibaba.alihealth.drug.kyt.saveent
//
// 新增往来单位企业记录
func AlibabaAlihealthDrugKytSaveent(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSaveentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSaveentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
