package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.synonymauths
//
// 物流企业查询货主企业信息
func AlibabaAlihealthDrugKytSynonymauths(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSynonymauthsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSynonymauthsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
