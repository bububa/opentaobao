package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcTopicGroupDelete 删除消息topic分组路由
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
func TaobaoTmcTopicGroupDelete(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcTopicGroupDeleteAPIRequest, resp *tmc.TaobaoTmcTopicGroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
