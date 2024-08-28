package kclub

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// AlibabaKclubKcQaGet 知识云-查询单个知识详情
// alibaba.kclub.kc.qa.get
//
// 知识云-查询单个知识详情。通过租户id、问题id查询问题详情
func AlibabaKclubKcQaGet(ctx context.Context, clt *core.SDKClient, req *kclub.AlibabaKclubKcQaGetAPIRequest, resp *kclub.AlibabaKclubKcQaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
