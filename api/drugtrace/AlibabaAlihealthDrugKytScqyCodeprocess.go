package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyCodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
func AlibabaAlihealthDrugKytScqyCodeprocess(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
