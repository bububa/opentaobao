package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanskusupplierquoteupload 标准供应商配额同步
// alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload
//
// 标准供应商配额同步
func Alibabatmallgeniescpplanskusupplierquoteupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanskusupplierquoteuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanskusupplierquoteuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanskusupplierquoteuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
