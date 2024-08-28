package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyQuerysubcodes 查询一个码的所有子码
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
func AlibabaAlihealthDrugKytYyQuerysubcodes(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
