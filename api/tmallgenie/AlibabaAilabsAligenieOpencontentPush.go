package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpencontentPush 天猫精灵内容接入标准接口
// alibaba.ailabs.aligenie.opencontent.push
//
// 第三方内容接入天猫精灵内容库，供相关技能使用
func AlibabaAilabsAligenieOpencontentPush(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpencontentPushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpencontentPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
