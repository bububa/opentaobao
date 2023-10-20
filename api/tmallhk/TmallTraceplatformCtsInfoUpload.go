package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmalltraceplatformctsinfoupload CTS提交溯源信息
// tmall.traceplatform.cts.info.upload
//
// cts上传溯源信息
func Tmalltraceplatformctsinfoupload(clt *core.SDKClient, req *tmallhk.TmalltraceplatformctsinfouploadAPIRequest, session string) (*tmallhk.TmalltraceplatformctsinfouploadAPIResponse, error) {
	var resp tmallhk.TmalltraceplatformctsinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
