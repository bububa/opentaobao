package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelCommentCreate 差评导入
// alibaba.wdk.channel.comment.create
//
// 差评导入
func AlibabaWdkChannelCommentCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelCommentCreateAPIRequest, resp *wdk.AlibabaWdkChannelCommentCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
