package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieVideoalbumPush 天猫精灵内容库视频合辑数据推送接口
// alibaba.ailabs.aligenie.videoalbum.push
//
// 三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
func AlibabaAilabsAligenieVideoalbumPush(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieVideoalbumPushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieVideoalbumPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
