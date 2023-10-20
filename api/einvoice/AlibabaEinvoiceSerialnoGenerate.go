package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceserialnogenerate 获取统一开票流水号
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
func Alibabaeinvoiceserialnogenerate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceserialnogenerateAPIRequest, session string) (*einvoice.AlibabaeinvoiceserialnogenerateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceserialnogenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
