package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanbomupload 计划BOM同步
// alibaba.tmallgenie.scp.plan.bom.upload
//
// 计划BOM同步
func Alibabatmallgeniescpplanbomupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanbomuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanbomuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanbomuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
