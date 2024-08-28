package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectChannelphone 新房渠道电话数据同步
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
func AlibabaAlihouseNewhomeProjectChannelphone(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
