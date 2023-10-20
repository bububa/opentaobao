package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlereportmediaupload 验货报告上传文件
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
func Alibabaidlereportmediaupload(clt *core.SDKClient, req *idle.AlibabaidlereportmediauploadAPIRequest, session string) (*idle.AlibabaidlereportmediauploadAPIResponse, error) {
	var resp idle.AlibabaidlereportmediauploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
