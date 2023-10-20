package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplancorrectsupplierquoteupload 供应商校准后的配额同步
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
func Alibabatmallgeniescpplancorrectsupplierquoteupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplancorrectsupplierquoteuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
