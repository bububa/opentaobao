package tax

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tax"
)

// Alibabataxinvoicesync 第三方开票回调API
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
func Alibabataxinvoicesync(clt *core.SDKClient, req *tax.AlibabataxinvoicesyncAPIRequest, session string) (*tax.AlibabataxinvoicesyncAPIResponse, error) {
	var resp tax.AlibabataxinvoicesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
