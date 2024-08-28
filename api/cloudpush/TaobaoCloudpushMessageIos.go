package cloudpush

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushMessageIos 百川云推送发送消息给ios
// taobao.cloudpush.message.ios
//
// 百川云推送发送消息给iOS设备.
func TaobaoCloudpushMessageIos(ctx context.Context, clt *core.SDKClient, req *cloudpush.TaobaoCloudpushMessageIosAPIRequest, resp *cloudpush.TaobaoCloudpushMessageIosAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
