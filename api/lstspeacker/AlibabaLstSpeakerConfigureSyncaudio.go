package lstspeacker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSyncaudio 音频同步
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
func AlibabaLstSpeakerConfigureSyncaudio(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
