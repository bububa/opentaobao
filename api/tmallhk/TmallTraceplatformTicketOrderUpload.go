package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmalltraceplatformticketorderupload 上传小票数据
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
func Tmalltraceplatformticketorderupload(clt *core.SDKClient, req *tmallhk.TmalltraceplatformticketorderuploadAPIRequest, session string) (*tmallhk.TmalltraceplatformticketorderuploadAPIResponse, error) {
	var resp tmallhk.TmalltraceplatformticketorderuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
