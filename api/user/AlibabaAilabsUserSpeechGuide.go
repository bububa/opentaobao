package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaAilabsUserSpeechGuide 引导语推荐接口
// alibaba.ailabs.user.speech.guide
//
// 根据用户的语音query，返回相应的引导语推荐
func AlibabaAilabsUserSpeechGuide(ctx context.Context, clt *core.SDKClient, req *user.AlibabaAilabsUserSpeechGuideAPIRequest, resp *user.AlibabaAilabsUserSpeechGuideAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
