package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// Alibabaalihealthmedicalimdataupload 三方IM图片音频消息上传
// alibaba.alihealth.medical.im.data.upload
//
// 三方IM图片音频消息上传
func Alibabaalihealthmedicalimdataupload(clt *core.SDKClient, req *alihealthmedical.AlibabaalihealthmedicalimdatauploadAPIRequest, session string) (*alihealthmedical.AlibabaalihealthmedicalimdatauploadAPIResponse, error) {
	var resp alihealthmedical.AlibabaalihealthmedicalimdatauploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
