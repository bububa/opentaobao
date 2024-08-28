package kclub

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// AlibabaKclubKcQueryknowledge 知识云-通用知识查询服务
// alibaba.kclub.kc.queryknowledge
//
// 知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
func AlibabaKclubKcQueryknowledge(ctx context.Context, clt *core.SDKClient, req *kclub.AlibabaKclubKcQueryknowledgeAPIRequest, resp *kclub.AlibabaKclubKcQueryknowledgeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
