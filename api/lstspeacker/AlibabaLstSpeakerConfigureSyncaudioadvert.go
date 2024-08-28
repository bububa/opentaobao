package lstspeacker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerConfigureSyncaudioadvert 同步广告
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
func AlibabaLstSpeakerConfigureSyncaudioadvert(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIRequest, resp *lstspeacker.AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
