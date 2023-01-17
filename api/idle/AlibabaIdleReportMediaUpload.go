package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleReportMediaUpload 验货报告上传文件
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
func AlibabaIdleReportMediaUpload(clt *core.SDKClient, req *idle.AlibabaIdleReportMediaUploadAPIRequest, session string) (*idle.AlibabaIdleReportMediaUploadAPIResponse, error) {
	var resp idle.AlibabaIdleReportMediaUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
