package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformAwdcInfoUpload AWDC提交溯源信息
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
func TmallTraceplatformAwdcInfoUpload(clt *core.SDKClient, req *tmallhk.TmallTraceplatformAwdcInfoUploadAPIRequest, session string) (*tmallhk.TmallTraceplatformAwdcInfoUploadAPIResponse, error) {
	var resp tmallhk.TmallTraceplatformAwdcInfoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
