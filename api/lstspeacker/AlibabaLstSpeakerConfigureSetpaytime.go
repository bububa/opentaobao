package lstspeacker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSetpaytime 音箱播放配置
// alibaba.lst.speaker.configure.setpaytime
//
// 音箱播放配置
func AlibabaLstSpeakerConfigureSetpaytime(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSetpaytimeAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureSetpaytimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
