package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerydruginfo 码查询药品
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
func AlibabaAlihealthDrugKytQuerydruginfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
