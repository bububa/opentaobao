package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.wes.synonymauths
//
// 物流企业查询货主企业信息
func AlibabaAlihealthDrugKytWesSynonymauths(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSynonymauthsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
