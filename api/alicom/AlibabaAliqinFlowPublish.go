package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowPublish 流量发放(用户id)
// alibaba.aliqin.flow.publish
//
// 阿里通信流量下发功能，允许用户补发
func AlibabaAliqinFlowPublish(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowPublishAPIRequest, resp *alicom.AlibabaAliqinFlowPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
