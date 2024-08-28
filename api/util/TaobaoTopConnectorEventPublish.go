package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopConnectorEventPublish 连接器事件发布
// taobao.top.connector.event.publish
//
// 连接器事件发布
func TaobaoTopConnectorEventPublish(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTopConnectorEventPublishAPIRequest, resp *util.TaobaoTopConnectorEventPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
