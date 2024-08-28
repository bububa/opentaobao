package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyPutpackage 码拼箱
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
func AlibabaAlihealthDrugKytScqyPutpackage(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
