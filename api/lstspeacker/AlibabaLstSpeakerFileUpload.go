package lstspeacker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerFileUpload 如意音箱音频文件长传
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
func AlibabaLstSpeakerFileUpload(ctx context.Context, clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerFileUploadAPIRequest, resp *lstspeacker.AlibabaLstSpeakerFileUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
