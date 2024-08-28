package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseQuerystandpointSave 法宝侧主动查询反馈
// alibaba.legal.case.querystandpoint.save
//
// 法宝侧主动查询反馈口径,此接口只用来反馈主动查询的口径,之前推送的口径反馈不适合
func AlibabaLegalCaseQuerystandpointSave(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseQuerystandpointSaveAPIRequest, resp *legalcase.AlibabaLegalCaseQuerystandpointSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
