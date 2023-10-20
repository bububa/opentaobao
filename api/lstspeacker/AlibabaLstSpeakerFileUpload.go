package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// AlibabaLstSpeakerFileUpload 如意音箱音频文件长传
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
func AlibabaLstSpeakerFileUpload(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerFileUploadAPIRequest, resp *lstspeacker.AlibabaLstSpeakerFileUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
