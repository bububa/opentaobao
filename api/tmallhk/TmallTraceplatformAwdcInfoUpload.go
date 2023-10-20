package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmalltraceplatformawdcinfoupload AWDC提交溯源信息
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
func Tmalltraceplatformawdcinfoupload(clt *core.SDKClient, req *tmallhk.TmalltraceplatformawdcinfouploadAPIRequest, session string) (*tmallhk.TmalltraceplatformawdcinfouploadAPIResponse, error) {
	var resp tmallhk.TmalltraceplatformawdcinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
