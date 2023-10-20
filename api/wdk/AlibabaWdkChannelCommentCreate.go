package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelCommentCreate 差评导入
// alibaba.wdk.channel.comment.create
//
// 差评导入
func AlibabaWdkChannelCommentCreate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelCommentCreateAPIRequest, resp *wdk.AlibabaWdkChannelCommentCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
