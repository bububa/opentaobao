package alidoc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthDrugdiseaseQuery 药品诊断查询
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
func AlibabaAlihealthDrugdiseaseQuery(ctx context.Context, clt *core.SDKClient, req *alidoc.AlibabaAlihealthDrugdiseaseQueryAPIRequest, resp *alidoc.AlibabaAlihealthDrugdiseaseQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
