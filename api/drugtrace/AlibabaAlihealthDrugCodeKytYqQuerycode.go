package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytYqQuerycode 查询追溯码对应的药品信息（疫情）
// alibaba.alihealth.drug.code.kyt.yq.querycode
//
// 通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。
func AlibabaAlihealthDrugCodeKytYqQuerycode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytYqQuerycodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytYqQuerycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
