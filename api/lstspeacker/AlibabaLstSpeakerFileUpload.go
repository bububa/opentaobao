package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerfileupload 如意音箱音频文件长传
// alibaba.lst.speaker.file.upload
//
// 如意音箱音频文件长传
func Alibabalstspeakerfileupload(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerfileuploadAPIRequest, session string) (*lstspeacker.AlibabalstspeakerfileuploadAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerfileuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
