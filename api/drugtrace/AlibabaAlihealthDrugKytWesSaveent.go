package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSaveent 新增往来单位企业记录
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
func AlibabaAlihealthDrugKytWesSaveent(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSaveentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSaveentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
