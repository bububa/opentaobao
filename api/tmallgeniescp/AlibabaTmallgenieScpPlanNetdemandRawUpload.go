package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplannetdemandrawupload 二级物料净需求回传(TL+1)
// alibaba.tmallgenie.scp.plan.netdemand.raw.upload
//
// 二级物料净需求回传(TL+1)
func Alibabatmallgeniescpplannetdemandrawupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplannetdemandrawuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplannetdemandrawuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplannetdemandrawuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
