package cloudpush

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushPush 百川用户使用云推送高级推送接口
// taobao.cloudpush.push
//
// 百川用户使用云推送高级推送接口
func TaobaoCloudpushPush(ctx context.Context, clt *core.SDKClient, req *cloudpush.TaobaoCloudpushPushAPIRequest, resp *cloudpush.TaobaoCloudpushPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
